package utils

import (
	"fmt"
	"github-telegram-notify/types"
	"html"
	"strings"
)

func CreateContents(meta *types.Metadata, authorTag string) (text string, markupText string, markupUrl string, err error) {
	event, _ := meta.ParseEvent()
	switch meta.EventName {
	case "pull_request":
		event := event.(*types.PullRequestEvent)

		if !Contains([]string{
			"created", "opened", "reopened", "locked", "unlocked", "closed", "synchronize", // More to be added.
		}, event.Action) {
			err = fmt.Errorf("unsupported event type '%s' for %s", event.Action, meta.EventName)
			return
		}

		text = createPullRequestText(event, authorTag)
		markupText = "Open Pull Request"
		markupUrl = event.PullRequest.HTMLURL
	case "pull_request_review_comment":
		event := event.(*types.PullRequestReviewCommentEvent)

		if !Contains([]string{"created", "deleted"}, event.Action) {
			err = fmt.Errorf("unsupported event type '%s' for %s", event.Action, meta.EventName)
			return
		}

		text = createPullRequestReviewCommentText(event, authorTag)
		markupText = "Open Comment"
		markupUrl = event.Comment.HTMLURL
	case "push":
		event := event.(*types.PushEvent)
		// No Activity Types
		text = createPushText(event, authorTag)
		markupText = "Open Changes"
		markupUrl = event.Compare
	case "release":
		event := event.(*types.ReleaseEvent)
		if !Contains([]string{"published", "released"}, event.Action) {
			err = fmt.Errorf("unsupported event type '%s' for %s", event.Action, meta.EventName)
			return
		}

		text = createReleaseText(event, authorTag)
		markupText = "🌐"
		markupUrl = event.Release.HTMLURL
	}
	return text, markupText, markupUrl, nil
}

func createPushText(event *types.PushEvent, authorTag string) string {
	text := fmt.Sprintf("<b>🔨 %d New commit to</b> <a href='%s'>%s</a>[<code>%s</code>]\n\n",
		len(event.Commits),
		event.Repo.HTMLURL,
		event.Repo.FullName,
		strings.Replace(event.Ref, "refs/heads/", "", 1),
	)

	for _, commit := range event.Commits {
		text += fmt.Sprintf("• <a href='%s'>%s</a> - %s by <a href='%s'>%s</a>\n",
			commit.Url,
			commit.Id[:7],
			html.EscapeString(commit.Message),
			commit.Author.HTMLURL,
			commit.Author.Name,
		)
	}

	if authorTag != "" {
		text += fmt.Sprintf("\n\n@%s, fyi", authorTag)
	}

	return text
}

func createPullRequestText(event *types.PullRequestEvent, authorTag string) (text string) {
	text = fmt.Sprintf("🔌 <a href='%s'>%s</a> ", event.Sender.HTMLURL, event.Sender.Login)
	text += event.Action
	if event.Action == "opened" {
		text += " a new"
	}
	text += " pull request "
	text += fmt.Sprintf("<a href='%s'>%s</a>", event.PullRequest.HTMLURL, html.EscapeString(event.PullRequest.Title))
	text += fmt.Sprintf(" in <a href='%s'>%s</a>", event.Repo.HTMLURL, event.Repo.FullName)

	if authorTag != "" {
		text += fmt.Sprintf("\n\n@%s, fyi", authorTag)
	}

	return text
}

func createPullRequestReviewCommentText(event *types.PullRequestReviewCommentEvent, authorTag string) string {
	text := fmt.Sprintf("🧐 <a href='%s'>%s</a> commented on PR review <a href='%s'>%s</a> in <a href='%s'>%s</a>",
		event.Sender.HTMLURL,
		event.Sender.Login,
		event.PullRequest.HTMLURL,
		html.EscapeString(event.PullRequest.Title),
		event.Repo.HTMLURL,
		event.Repo.FullName,
	)

	if authorTag != "" {
		text += fmt.Sprintf("\n\n@%s, fyi", authorTag)
	}

	return text
}

func createReleaseText(event *types.ReleaseEvent, authorTag string) (text string) {
	text = "🎊 A new "
	if event.Release.Prerelease {
		text += "pre"
	}
	text += fmt.Sprintf("release was %s in <a href='%s'>%s</a> by <a href='%s'>%s</a>\n",
		event.Action,
		event.Repo.HTMLURL,
		event.Repo.FullName,
		event.Sender.HTMLURL,
		event.Sender.Login,
	)
	text += fmt.Sprintf("\n📍 <a href='%s'>%s</a> (<code>%s</code>)\n\n", event.Release.HTMLURL, event.Release.Name, event.Release.TagName)
	if event.Release.Assets != nil {
		text += "📦 <b>Assets:</b>\n"
		for _, asset := range event.Release.Assets {
			text += fmt.Sprintf("• <a href='%s'>%s</a>\n", asset.BrowserDownloadURL, html.EscapeString(asset.Name))
		}
	}

	if authorTag != "" {
		text += fmt.Sprintf("\n\n@%s, fyi", authorTag)
	}

	return
}