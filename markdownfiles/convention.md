## Code

- For private variable and function/method follow ```Camel case``` convention.
- For public variable and function/method follow ```Pascal case``` convention.
- Write unit testable code.
- Create and manage test fies in the same package.
- Add testcases considering corner cases.
- Please follow interface first development.

## Commit

- Please use proper commit message. Example:
    - ```Add Github webhook integration``` or
    - ```Enable Github webhook integration```
- In body, please explain , ```What is the problem it is going to solve!```. Try to add a view of the interface.
  Example:

```
Creates webhook to listen repository code change events.  

type Git interface {
    ...
    CreateRepositoryWebhook(username,repogitory_name,token string)(v1.GithubWebhook,error)
    ...
}
```
