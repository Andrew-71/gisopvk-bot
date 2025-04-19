package mock

import (
	"strings"

	"git.a71.su/Andrew71/gisopvk-bot/internal/domain"
)

var RegLink = "https://gisopvk.ru/media/files/filepublic/d/a/a/daa6eeceb30a1625.как-прописать-права-доступа-сотруднику.pdf"
var SignLink = "https://gisopvk.ru/media/files/filepublic/8/d/8/8d8c960b707c8b38.памятка-эп_1.pdf"

type MockBot struct {
}

func (b MockBot) Reply(message domain.Query) (domain.Reply, error) {
	// this is horrific. straight up horrific.
	body := strings.ToLower(message.Body)
	if strings.Contains(body, "регистр") {
		return domain.Reply{message.UUID, "По вопросам регистрации, пожалуйста обращайтесь к данному файлу - " + RegLink}, nil
	} else if strings.Contains(body, "эп") || strings.Contains(body, "рабочее место") {
		return domain.Reply{message.UUID, "По вопросам настройки рабочего места для работы с ЭП, пожалуйста обращайтесь к данному файлу - " + SignLink}, nil
	}
	return domain.Reply{message.UUID,
		"Я не уверен, что понимаю ваш вопрос. Вы можете спросить меня о регистрации или настройке рабочего места для работы с ЭП"}, nil
}

func NewMockBot() MockBot {
	return MockBot{} // has no state... for now
}
