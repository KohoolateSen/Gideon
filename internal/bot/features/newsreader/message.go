package newsreader

import (
	"fmt"
	"strings"
)

func (N *NewsReader) FormatArticle(article Article, index int) string {
	var builder strings.Builder
	addLine := func(label, value string) {
		if value != "" {
			builder.WriteString("\n" + label + value + "\n")
		}
	}
	addLine("🏷️ 𝐓𝐢𝐭𝐥𝐞 : ", article.Title)
	addLine("📅 𝐏𝐮𝐛𝐥𝐢𝐬𝐡𝐞𝐝 𝐚𝐭 : ", article.PublishedAt)
	addLine("📝 𝐃𝐞𝐬𝐜𝐫𝐢𝐩𝐭𝐢𝐨𝐧 : ", article.Description)
	addLine("🔗 𝐔𝐑𝐋 : ", article.URL)
	addLine("📖 𝐑𝐞𝐬𝐮𝐥𝐭 : ", fmt.Sprintf("%d 𝐟𝐫𝐨𝐦 %d", index, len(N.Response.Articles)-1))
	return builder.String()
}

func (N *NewsReader) WrapArticleAsMessage() string {
	article := N.Response.Articles[N.Cpage%len(N.Response.Articles)]
	return N.FormatArticle(article, N.Cpage)
}

func (N *NewsReader) WrapArticleAsMessageByIndex(index int) string {
	article := N.Response.Articles[index]
	return N.FormatArticle(article, index)
}
