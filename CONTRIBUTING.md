# How to become a contributor

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

✔️ You agree to the [Contributor License Agreements](https://www.contributor-covenant.org)

✔️ Make sure your code adheres to the existing style, technologies, and coding standards recommended for this project.

### Did you find a bug? / Do you want to suggest something?

- Ensure the bug was not already reported by searching on [Issues](https://github.com/EpitechMscProPromo2025/T-YEP-600-NAN-6-1-finalproject-ange.marchand/issues).  
  Otherwise, create an new one. Be sure to include a **clear title and description**, as much relevant information as possible, and a **sample code** or **executable test case** demonstrating the expected behavior that does not not occur.

- Use issue creation templates

### **Do you want to create a branch?**

- Your branch name should be formatted as `fix/<ISSUENUMBER>-<TITLE>` for bug fixes or `feature/<ISSUENUMBER>-<TITLE>` for features.  
  Example:  
  `fix/4221-infinite-loop`  
  `feature/4222-deployement`  
  `doc/4223-RGPD`

### **Do you want to fix an issue?**

- Create a new branch following the above convention
- Implement your features of fixes in it.

### **How to title commits?**

- Commit often!

- Follow [the guidelines](https://cbea.ms/git-commit/)

- Use imperative tense (avoid past tense).

- The title of the commit must be a summuary of the content and not be too long (less than 50 characters).

- Prefer putting detailed informations inside the commit's description.

  ```sh
  Example:

  $> git commit -m 'Infinite loop when pressing Alt-F4

  This was caused by a missing check in the event loop
  The program now checks when the window is set to close'
  ```

### How to submit a pull request?

- Format your code using Prettier

```sh
yarn run format
yarn run lint
```

- Make sure your code has a proper set of unit tests that all pass

```sh
yarn run test:unit
```

- Once validated, merge to PR to `main` and remove the source branch (with `git branch -D <branch_name>`.

---

THANKS! 💚

VirtuozPlay team
