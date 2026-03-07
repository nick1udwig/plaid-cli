---
title: "Link - Handling an invalid Link Token | Plaid Docs"
source_url: "https://plaid.com/docs/link/handle-invalid-link-token/"
scraped_at: "2026-03-07T22:05:04+00:00"
---

# Handling an invalid Link Token

#### Catch the error in the onExit callback and refetch a new link\_token for the next time the user opens Link

Occasionally, the end user may invalidate the existing `link_token` that was used to open Link by taking too long to go through the flow (30+ minutes), or attempting too many invalid logins. If this happens, Link will exit with an [`INVALID_LINK_TOKEN`](/docs/errors/invalid-input/#invalid_link_token) error code.

To allow your user to open Link again, recognize the error in the `onExit` callback, fetch a new `link_token`, and use it to reinitialize Link. You can obtain a new `link_token` by making another [`/link/token/create`](/docs/api/link/#linktokencreate) request:

```
app.post('/api/create_link_token', async function (request, response) {
  // Get the client_user_id by searching for the current user
  const user = await User.find(...);
  const clientUserId = user.id;
  const linkTokenRequest = {
    user: {
      // This should correspond to a unique id for the current user.
      client_user_id: clientUserId,
    },
    client_name: 'Plaid Test App',
    products: ['auth'],
    language: 'en',
    webhook: 'https://webhook.example.com',
    redirect_uri: 'https://domainname.com/oauth-page.html',
    country_codes: ['US'],
  };
  try {
    const createTokenResponse = await client.linkTokenCreate(linkTokenRequest);
    response.json(createTokenResponse.data);
  } catch (error) {
    // handle error
  }
});
```

For the Link web integration, reinitializing Link means creating a new iframe. To avoid stacking iframes for each Link initialization, you can clean up the old iframe by calling the [`destroy()`](/docs/link/web/#destroy) method on the Plaid Link handler.

```
// Initialize Link with a new link_token each time.
const configs = {
  token: (await $.post('/create_link_token')).link_token,
  onSuccess: (public_token, metadata) => {
    // Send the public_token to your app server.
  },
  onExit: (err, metadata) => {
    // The user exited the Link flow with an INVALID_LINK_TOKEN error.
    // This can happen if the token expires or the user has attempted
    // too many invalid logins.
    if (err != null && err.error_code === 'INVALID_LINK_TOKEN') {
      linkHandler.destroy();
      linkHandler = Plaid.create({
        ...configs,
        // Fetch a new link_token because the old one was invalidated.
        token: (await $.post('/create_link_token')).link_token,
      });
    }
    // metadata contains the most recent API request ID and the
    // Link session ID. Storing this information is helpful
    // for support.
  },
};

let linkHandler = Plaid.create(configs);
```

When the user is ready, they will be able to reopen Link and go through the authentication process again.

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
