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

