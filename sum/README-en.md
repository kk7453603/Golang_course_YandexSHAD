[Русский](README.md)
# sum

In this task, you need to learn how to submit solutions to the evaluation system.

0. Clone this repo https://gitlab.com/slon/shad-go

1. Open file `sum.go` and implement a function that returns the sum of two numbers.

2. Make sure that your solution passes the tests locally

   ```shell
   # run from the root of the repo
   go test ./sum/...
   ```
   
3. Make sure that your code passes the linter. [Linter setup](https://github.com/golangci/golangci-lint#binary).

   ```shell
   # run from the root of the repo
   golangci-lint run ./sum/...
   ```

4. **(Do it only once)** Setup remote for you private repo.

   ```shell
   # Sign up in the test system if you haven't done it already.
   # Click on "My Repo" link at https://go.manytask.org/
   # Press the blue Clone button and copy the address from "Clone with SSH" field
   # Run the following  command in shell, replacing the last argument with your copied repo address
   git remote add student git@gitlab.manytask.org:go-spring-2023/USERNAME.git
   ```

5. **(Do it only once)** [Setup](https://gitlab.manytask.org/-/profile/keys) ssh key. If you're not familiar with it please follow
   the instructions on the linked page.

6. Add and commit your changes to git.

   ```shell
   git add .
   git commit -m "Solved sum"
   ```
   
7. Push your commits to remote branch.

   ```shell
   git push student
   ```
   
   **NOTE:** The system checks only tasks that were changed in the last pushed commit. If you made multiple commits
   and did a single push, only the last commit will be checked. If you want to restart checks for a specific commit
   you can press "Retry" on the page with check logs or push a new commit with minor changes.
   
8. You can check at task evaluation process by clicking on the "Submit" link on the page https://go.manytask.org/

9. Make sure that your grade has been added to [the table](https://docs.google.com/spreadsheets/d/1rTqdHu2AJtdCeFp_iLKsKVk848Wddh2rW0nOmO6Y528).
