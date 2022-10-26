1. Understanding how go-plugin works
2. Implementing an example plugin to have it implemented myself
3. Basic mechanism to communicate between Gitea and plugins, including data model, and passing the most important information from Gitea to the plugin such as the absolute instance URL, or Gitea version, or a way to perform an API call
4. Custom routes
5. Custom tabs
6. Hooks for certain system events, i.e. plugin initialization, application startup or shutdown
7. Communication between plugins
8. custom resource strings for translations
9. Configuration for enabling plugins
10. global limits for plugins, i.e. how long a plugin call may take at most before being cancelled
11. Document how to add a new plugin
12. Document how the plugin mechanism works internally
