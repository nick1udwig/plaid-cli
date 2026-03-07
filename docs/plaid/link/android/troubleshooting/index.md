---
title: "Link - Troubleshooting | Plaid Docs"
source_url: "https://plaid.com/docs/link/android/troubleshooting/"
scraped_at: "2026-03-07T22:05:02+00:00"
---

# Troubleshooting the Plaid Link Android SDK

=\*=\*=\*=

#### Enabling Logs

The Link SDK logs information to LogCat at several points in the flow. Pass a `LinkLogLevel` value with the
`LinkTokenConfiguration` to see the logs. The levels from least to most verbose: `ERROR`, `WARN`, `INFO`, `DEBUG`, `VERBOSE`.

openLink

```
val linkTokenConfiguration = linkTokenConfiguration {
  token = "LINK_TOKEN_FROM_SERVER"
  logLevel = if (BuildConfig.DEBUG) LinkLogLevel.VERBOSE else LinkLogLevel.ERROR
}
```

=\*=\*=\*=

#### Troubleshooting OAuth errors

A troubleshooting guide for common OAuth errors is below.

=\*=\*=\*=

#### No redirect out of app

##### Link user experience

- In Link, after clicking "Continue" on the OAuth screen, nothing happens.

##### Common causes

- The user may be using an unsupported browser that is not compatible with Plaid's OAuth redirects, such as DuckDuckGo. For more details on which browsers are supported, see [Supported browsers](https://plaid.com/docs/link/web/#supported-browsers).

##### Troubleshooting steps

Try again with Google Chrome.

=\*=\*=\*=

#### No redirect back to app

##### Link user experience

- After completing OAuth in the browser or financial institution's app, the user is not redirected back to your app.

##### Common causes

- The webpage is unable to locate an app on the device with a package id matching the one used to create the Link token.

##### Troubleshooting steps

Check that the package id you passed in matched the package id of the app you are using.

=\*=\*=\*=

#### The Play Store opened upon redirect

##### Link user experience

- After completing OAuth in the browser or financial institution's app, the Google Play store opened.

##### Common causes

- The webpage is unable to locate an app on the device with the package id matching the one used to create the Link token.

##### Troubleshooting steps

Check that the package name used to create the Link token matches the package name of the app it was used in.

=\*=\*=\*=

#### Link opens, then immediately closes upon redirect

##### Link user experience

After completing OAuth in the browser or financial institution's app, Link opens and then closes again.

##### Common causes

- The webpage redirected back to the wrong application on the device causing Link to open and immediately close again as it gets data from a different session.
- You may have both a test and a release version of your app installed on your device and used the wrong package name when creating the Link token.

##### Troubleshooting steps

Check that the package name used to create the Link token matches the package name of the app it was used in.

=\*=\*=\*=

#### Other common errors

To troubleshoot an error with an error code, use the [error troubleshooting guide](/docs/errors/).

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
