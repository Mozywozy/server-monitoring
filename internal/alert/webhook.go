package alert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// WebhookPayload digunakan untuk memformat payload JSON
type WebhookPayload struct {
	Content string `json:"content"` // Menggunakan "content" sebagai kunci
}

// SendWebhook mengirimkan pesan ke URL webhook yang ditentukan
func SendWebhook(webhookURL, message string) error {
	client := &http.Client{}

	// Membuat payload JSON
	payload := WebhookPayload{
		Content: message, // Menggunakan Content
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err
	}

	fmt.Println("Sending webhook to:", webhookURL)
	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	defer response.Body.Close()

	// Membaca body respons
	if response.StatusCode != http.StatusOK {
		responseBody, _ := ioutil.ReadAll(response.Body)
		fmt.Printf("Failed to send webhook, status code: %d, response: %s\n", response.StatusCode, string(responseBody))
		return fmt.Errorf("failed to send webhook, status code: %d", response.StatusCode)
	}

	fmt.Println("Webhook sent successfully")
	return nil
}
