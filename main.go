package firebase_messaging

import (
        "context"
        "fmt"

        firebase "firebase.google.com/go"
        "firebase.google.com/go/messaging"
        "google.golang.org/api/option"
)

func Send(pathToServiceAccountKey string, msg *messaging.Message) (string, error) {
        client, err := getClient(pathToServiceAccountKey)
        if err != nil {
                return "", err
        }

        response, err := client.Send(context.Background(), msg)
        if err != nil {
                return "", wrappedError("Error sending message via client", err)
        }

        return response, nil
}

func getClient(pathToServiceAccountKey string) (*messaging.Client, error) {
        opt := option.WithCredentialsFile(pathToServiceAccountKey)
        app, err := firebase.NewApp(context.Background(), nil, opt)
        if err != nil {
                return nil, wrappedError("Error initializing push notifications", err)
        }

        // Obtain a messaging.Client from the App.
        ctx := context.Background()
        client, err := app.Messaging(ctx)
        if err != nil {
                return nil, wrappedError("Error creating message client", err)
        }

        return client, nil
}

func wrappedError(msg string, err error) (error) {
        return fmt.Errorf("%s: %v", msg, err)
}
