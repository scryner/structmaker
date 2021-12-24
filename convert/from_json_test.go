package convert

import "testing"

const testSimple = `{
	"dog_name": "bokdol",
	"dog_age": 5,
	"owner": {
		"name": "scryner"
	}
}`

const testComplex = `{
  "after": "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
  "before": "95790bf891e76fee5e1747ab589903a6a1f80f22",
  "checkout_sha": "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
  "commits": [
    {
      "author": {
        "email": "test@example.com",
        "name": "Test User"
      },
      "id": "c5feabde2d8cd023215af4d2ceeb7a64839fc428",
      "message": "Add simple search to projects in public area\n\ncommit message body",
      "timestamp": "2013-05-13T18:18:08+00:00",
      "title": "Add simple search to projects in public area",
      "url": "https://test.example.com/gitlab/gitlab/-/commit/c5feabde2d8cd023215af4d2ceeb7a64839fc428"
    }
  ],
  "event_name": "push",
  "message": "Hello World",
  "object_kind": "push",
  "project": {
    "avatar_url": "https://s.gravatar.com/avatar/d4c74594d841139328695756648b6bd6?s=8://s.gravatar.com/avatar/d4c74594d841139328695756648b6bd6?s=80",
    "default_branch": "master",
    "description": "",
    "git_http_url": "http://test.example.com/gitlab/gitlab.git",
    "git_ssh_url": "git@test.example.com:gitlab/gitlab.git",
    "id": 15,
    "name": "gitlab",
    "namespace": "gitlab",
    "path_with_namespace": "gitlab/gitlab",
    "visibility_level": 0,
    "web_url": "http://test.example.com/gitlab/gitlab"
  },
  "project_id": 15,
  "push_options": {
    "ci": {
      "skip": true
    }
  },
  "ref": "refs/heads/master",
  "total_commits_count": 1,
  "user_avatar": "https://s.gravatar.com/avatar/d4c74594d841139328695756648b6bd6?s=8://s.gravatar.com/avatar/d4c74594d841139328695756648b6bd6?s=80",
  "user_email": "john@example.com",
  "user_id": 4,
  "user_name": "John Smith"
}`

func TestFromJson(t *testing.T) {
	// simple
	t.Run("simple", testConvert(testSimple, "Dog"))

	// complex
	t.Run("complex", testConvert(testComplex, "TagEvent"))
}

func testConvert(content, baseName string) func(*testing.T) {
	return func(t *testing.T) {
		s, err := FromJson([]byte(content), baseName)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		t.Log(s)
	}
}
