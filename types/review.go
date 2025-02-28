package types

/*
review object Required
The review that was affected.

Properties of review
Name, Type, Description
_links object Required
Properties of _links
author_association string Required
How the author is associated with the repository.

Can be one of: COLLABORATOR, CONTRIBUTOR, FIRST_TIMER, FIRST_TIME_CONTRIBUTOR, MANNEQUIN, MEMBER, NONE, OWNER

body string or null Required
The text of the review.

commit_id string Required
A commit SHA for the review.

html_url string Required
id integer Required
Unique identifier of the review

node_id string Required
pull_request_url string Required
state string Required
Can be one of: dismissed, approved, changes_requested

submitted_at string Required
user object or null Required
*/

type Review struct {
	// AuthorAssociation Can be one of: COLLABORATOR, CONTRIBUTOR, FIRST_TIMER, FIRST_TIME_CONTRIBUTOR, MANNEQUIN, MEMBER, NONE, OWNER
	AuthorAssociation string `json:"author_association,omitempty"`
	Body              string `json:"body,omitempty"`
	CommitID          string `json:"commit_id,omitempty"`
	HTMLURL           string `json:"html_url,omitempty"`
	ID                int    `json:"id,omitempty"`
	NodeID            string `json:"node_id,omitempty"`
	PullRequestURL    string `json:"pull_request_url,omitempty"`
	// State Can be one of: dismissed, approved, changes_requested
	State       string `json:"state,omitempty"`
	SubmittedAt string `json:"submitted_at,omitempty"`
	User        User   `json:"user,omitempty"`
}