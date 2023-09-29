package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/Ahmed-Sermani/driver-tracking-system/users/usersclient"
	"github.com/Ahmed-Sermani/driver-tracking-system/users/usersclient/models"
	"github.com/Ahmed-Sermani/driver-tracking-system/users/usersclient/user"
	"github.com/centrifugal/centrifuge-go"
)

var (
	usersHost       = flag.String("users-host", "127.0.0.1", "the host of the users service")
	wsHost          = flag.String("ws-host", "127.0.0.1", "the host of the messaging service")
	usersBase       = flag.String("users-base", "/app", "the base url for users service")
	wsBase          = flag.String("ws-base", "/messaging", "the base url for messaging service")
	numGor          = flag.Int("num-gor", 1, "the number of goroutines to spawn")
	logPublications = false
)

func main() {
	flag.Parse()

	wg := sync.WaitGroup{}
	wg.Add(*numGor)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 0; i < *numGor; i++ {
		go func() {
			defer wg.Done()
			run(ctx)
		}()
		if i%100 == 0 {
			time.Sleep(3 * time.Second)
		}
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	select {
	case <-sigCh:
		cancel()
	case <-ctx.Done():
	}

	wg.Wait()
}

func run(ctx context.Context) {
	users := usersclient.NewHTTPClientWithConfig(
		nil,
		&usersclient.TransportConfig{
			Host:     *usersHost,
			BasePath: *usersBase,
		},
	)

	username := fmt.Sprintf("testuser-%d", rand.Intn(9999999))
	phone := fmt.Sprintf("%d", rand.Intn(9999999999))
	resp, err := users.User.RegisterUser(
		&user.RegisterUserParams{
			Body: &models.UserspbRegisterUserRequest{
				Name:  username,
				Phone: phone,
			},
			Context: context.Background(),
		},
	)
	if err != nil {
		log.Printf("failed %s", err.Error())
		return
	}
	if !resp.IsSuccess() {
		log.Printf("failed %s", resp.Error())
		return
	}
	userID := resp.GetPayload().ID

	loginResp, err := users.User.Login(
		&user.LoginParams{
			Body: &models.UserspbLoginRequest{
				Phone: phone,
			},
			Context: context.Background(),
		},
	)
	if err != nil {
		log.Printf("failed %s", err.Error())
		return
	}
	if !loginResp.IsSuccess() {
		log.Printf("failed %s", loginResp.Error())
		return
	}
	token := loginResp.GetPayload().Token
	headers := http.Header{}
	headers.Add("Authorization", fmt.Sprintf("bearer %s", token))

	client := centrifuge.NewJsonClient(
		fmt.Sprintf("ws://%s%s/connection/websocket", *wsHost, *wsBase),
		centrifuge.Config{
			Header: headers,
		},
	)

	client.OnConnected(func(e centrifuge.ConnectedEvent) {
		fmt.Println("Connected to server")
	})

	client.OnDisconnected(func(e centrifuge.DisconnectedEvent) {
		fmt.Println("Disconnected from server")
	})

	client.OnError(func(e centrifuge.ErrorEvent) {
		fmt.Println("Error: ", e.Error.Error())
	})

	err = client.Connect()
	if err != nil {
		log.Print(err)
		return
	}

	defer client.Close()

	sub, err := client.NewSubscription(fmt.Sprintf("deliveries:%s", userID))
	if err != nil {
		log.Print(err)
		return
	}

	sub.OnSubscribed(func(e centrifuge.SubscribedEvent) {
		fmt.Println("Subscribed to channel: ", sub.Channel)
	})

	sub.OnPublication(func(pe centrifuge.PublicationEvent) {
		if logPublications {
			fmt.Println("Publication: ", string(pe.Data))
		}
	})

	err = sub.Subscribe()
	if err != nil {
		log.Print(err)
		return
	}

	defer sub.Unsubscribe()

	select {
	case <-ctx.Done():
	}
}
