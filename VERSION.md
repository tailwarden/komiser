# Komiser Versioning Guidelines

We use the [**SemVer**](https://semver.org/) versioning and maintain our own We follow the guidelines of Semantic Versioning (SemVer) and maintain our own standards on top of it.

## For a Komiser version x.y.z:

> X: Major | Y: Minor | Z: Patch

### Patch releases include

Patch releases incorporate performance and UX improving features without affecting the state in the forms of:

- Minute bug patches
- Cost coverage enhancements
- Supporting new cloud resources in the release that are backward compatible
- Any changes do **NOT** involve dealing with
  - Persistent state
  - Database or model changes

### Minor releases include

Minor releases incorporate significant features that maintain backward compatibility by:

- Keeping with the persistent state
- Handling logic for new as well old cases
- Database changes but have auto migrations that work with old models as well

### Major releases include

Major releases incorporate significant changes that fall into two main categories:

- Breaking change that is **NOT** backward compatible
- Huge enhancements that might require a lot of refactoring

> ⛔ We prioritize simplicity in our versioning approach and generally avoid the use of other somewhat complicated versioning labels such as `alpha, beta, and rc`.
>
> However, if necessary, we are open to utilizing these labels as well to ensure the most appropriate versioning for our releases.

### Release Schedules

Our ideal release schedule aims to have regular releases once every 2 weeks. The version format will follow the guidelines mentioned above, incorporating the principles of Semantic Versioning (SemVer) and our own additional standards.

This systematic approach ensures clarity and consistency in our versioning process, making it easier for users and enterprises to understand the significance of each release and determine when to upgrade.
