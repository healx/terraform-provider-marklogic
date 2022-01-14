resource "marklogic_user" "test-user" {
  name     = "test-user"
  roles    = ["manage-user"]
  password = "foobar1234"
}
