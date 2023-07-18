# `Jiffy` 🤖

`Jiffy` is a small helpful robot that automates laborious and time-consuming manual tasks of collecting critical CVE (Common Vulnerabilities and Exposures) information required for filing Red Hat Security Advisories (RHSA). Named after its ability to perform tasks swiftly, Jiffy streamlines the process saving you valuable time and effort.

## The Motivation and Concept

The main motivation for bringing jiffy to life was the need for auto-curating CVE details for filing Red Hat Security Advisories (RHSA). It is quite a laborious manual process to collect all the CVE information required to file a RHSA advisory. The information required to fill out all the fields is scattered between Jira and Bugzilla. What a hassle to have to:
- Evaluate what CVEs have been fixed between two releases
- What types of problems they were
- Corresponding Jira and bugzilla id
- Order the CVEs by priority of impact (first critical, then important, then moderate, then low).
- Identify CVEs of the same priority and order them numerically by ID.

Wouldn't it be great to get all this information compiled for you automatically? Let jiffy take care of that for you. It automates this process and curates all the required details in a specific format to file RHSA advisory.

## Setup

First, generate a JIRA token by following the below steps:

- Navigate to the [jira application](https://issues.redhat.com/) and login with your sso credentials.
- In Jira, select your profile picture at the top right of the screen, then choose **Profile**. Once you access your profile, select **Personal Access Tokens** in the left-hand menu.
- Select **Create token**. Give your new token a name. Optionally, for security reasons, you can set your token to automatically expire after a set number of days.
- Click **Create**.

> **Note:** Permission level of personal access tokens is set to the level of access you currently have.

Secondly, you need to pass a list of the CVEs that have been fixed for a particular release. This can be done by [creating a filter in Jira](https://jexo.io/blog/how-to-create-filter-jira/), aka [JQL(Jira Query Language)](https://support.atlassian.com/jira-service-management-cloud/docs/use-advanced-search-with-jira-query-language-jql/). For example, I create the below filter that checks the following:

> project = IPT AND status in ("Selected for Development") AND labels = security ORDER BY status DESC, created DESC, duedate

- `project = IPT`: This part filters the issues to only those belonging to the project with the key "IPT". It restricts the search to a specific project in Jira.

- `status in ("Selected for Development")`: This part narrows down the search to issues that have the status "Selected for Development." The in operator allows searching for multiple statuses if required.

- `labels = security`: This section filters the issues further, selecting only those that have a label "security". Labels in Jira help categorize issues and make it easier to find and manage them based on their characteristics.

- `ORDER BY status DESC, created DESC, duedate`: This part specifies the order in which the search results should be displayed. The ORDER BY clause allows you to sort the issues based on specific fields. In this case, the results are sorted in descending order (DESC) for three fields: status, created, and duedate.

## Usage

Simply update the generated jira token and jql in the `.env` file and run the below command:

> go run main.go

## Future Prospects

Thank you for your interest in our project! 

This project is currently a work in progress (WIP), and we are in the process of transforming it into a robust GO CLI tool, powered by the [Cobra library](https://pkg.go.dev/github.com/spf13/cobra). Our primary goal is to deliver a seamless and user-friendly command-line experience, making installation on your systems effortless.

Excitingly, we also have plans to merge this project with [Apicurio CVE CLI](https://github.com/Apicurio/apicurio-cve-cli) to create a comprehensive CLI tool for managing Apicurio Product Releases.

Stay tuned for updates!

## Acknowledgments

We sincerely thank everyone who has contributed to the development of Jiffy. Your support and efforts are greatly appreciated.
