package reminai

import (
	"context"
	"fmt"
	"time"

	"github.com/sashabaranov/go-openai"
	"github.com/sawalreverr/recything/config"
	cdt "github.com/sawalreverr/recything/internal/custom-data"
	rai "github.com/sawalreverr/recything/internal/remin-ai"
)

var (
	FirstInstructions = "Anda adalah seorang ahli dalam bidang Recything, sebuah aplikasi yang bertujuan untuk memfasilitasi masyarakat dengan menyediakan akses untuk mencari tahu tentang apa dan bagaimana sampah dapat digunakan, berkompetisi untuk mendapatkan penghargaan, dan melaporkan jika mereka melihat sampah yang tidak pada tempatnya. Berikut adalah dataset yang akan Anda gunakan sebagai referensi:"

	LastInstructions = "Berikan jawaban dalam bentuk paragraf dengan maksimal 2 paragraf yang menjelaskan secara lengkap terkait pertanyaan pengguna. Jika pertanyaan tidak berhubungan dengan topik-topik di atas, jawab dengan 'Saya tidak tahu jawabannya karena pertanyaan Anda tidak terkait dengan aplikasi kami'."
)

type reminAIUsecase struct {
	customDataRepository cdt.CustomDataRepository
}

func NewReminAIUsecase(repo cdt.CustomDataRepository) rai.ReminAIUsecase {
	return &reminAIUsecase{customDataRepository: repo}
}

func (uc *reminAIUsecase) AskGPT(question rai.RequestInput) (string, error) {
	var dataset string
	datas, _, _ := uc.customDataRepository.FindAll(1, 1000, "created_at", "asc")

	apikey := config.GetConfig().OpenAI.APIKey
	client := openai.NewClient(apikey)

	for i, data := range *datas {
		dataset += fmt.Sprintf("%d. %v\nDeskripsi: %v\n\n", i+1, data.Topic, data.Description)
	}

	fullInstructions := fmt.Sprintf("%v\n%v\n%v", FirstInstructions, dataset, LastInstructions)

	assistantName := "ReMin AI"
	assistantRequest := openai.AssistantRequest{
		Name:         &assistantName,
		Model:        openai.GPT3Dot5Turbo,
		Instructions: &fullInstructions,
	}

	assistant, err := client.CreateAssistant(context.Background(), assistantRequest)
	if err != nil {
		return "", err
	}

	thread, err := client.CreateThread(context.Background(), openai.ThreadRequest{})
	if err != nil {
		return "", err
	}

	messageRequest := openai.MessageRequest{
		Role:    string(openai.ThreadMessageRoleUser),
		Content: question.Question,
	}
	_, err = client.CreateMessage(context.Background(), thread.ID, messageRequest)
	if err != nil {
		return "", err
	}

	run, err := client.CreateRun(context.Background(), thread.ID, openai.RunRequest{AssistantID: assistant.ID})
	if err != nil {
		return "", err
	}

	for run.Status != openai.RunStatusCompleted {
		time.Sleep(5 * time.Second)
		run, err = client.RetrieveRun(context.Background(), thread.ID, run.ID)
		if err != nil {
			return "", err
		}
	}

	msgs, err := client.ListMessage(context.Background(), thread.ID, nil, nil, nil, nil)
	if err != nil {
		return "", err
	}
	msg := msgs.Messages[0]

	return msg.Content[0].Text.Value, nil
}
