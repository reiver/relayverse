# relayverse

**relayverse** is an actor caching server for decentralized social-media (DeSo), including the Fediverse.

## URLs

There are 3 important URLs for **relayverse**:

* `/.well-known/acct-cache?resource=???`
* `/.well-known/acct-icon?resource=???`
* `/.well-known/acct-image?resource=???`

Where `???` is replace by a acct-URI.

For example:

* `https://example.com/.well-known/acct-cache?resource=acct:reiver@mastodon.social`
* `https://example.com/.well-known/acct-icon?resource=acct:reiver@mastodon.social`
* `https://example.com/.well-known/acct-image?resource=acct:reiver@mastodon.social`

### acct-cache

The `acct-cache` style URL returns the (cached) activity-JSON (`application/activity+json`) file for the account.

### acct-icon

The `acct-icon` style URL redirects to the icon-url for the account, as provided in (cached) activity-JSON (`application/activity+json`) file.

You can put this type of URL into an HTML `<img>` element.
For example:

```html
<img src="https://example.com/.well-known/acct-icon?resource=acct:reiver@mastodon.social" />
```

### acct-image

The `acct-image` style URL redirects to the image-url for the account, as provided in (cached) activity-JSON (`application/activity+json`) file.

You can put this type of URL into an HTML `<img>` element.
For example:

```html
<img src="https://example.com/.well-known/acct-image?resource=acct:reiver@mastodon.social" />
```

## Fediverse ID

A **Fediverse ID** (i.e., **Fediverse Address**), such as:

`@reiver@mastodon.social`

Can be turned into an acct-URI — by replacing the first `@` with `acct:`, as in:

`acct:reiver@mastodon.social`

## Configuration

There are 3 environment variables that can be used to configure **relayverse**:

* `CACHE_GC_SLEEP`
* `CACHE_MIN`
* `PORT`

### `CACHE_GC_SLEEP`

`CACHE_GC_SLEEP` sets the minimum duration that the cache garbage-collector sleeps for before it cleans-up again.
For example:

`CACHE_GC_SLEEP=1h20m30s relayverse`

### `CACHE_MIN`

`CACHE_MIN` sets the minimum duration that things are cached for.
For example:

`CACHE_MIN=10m20s relayverse`

### `PORT`

`PORT` sets the TCP-port used by the HTTP server.
For example:

`PORT=9090 relayverse`

## Author

Software **relayverse** was written by [Charles Iliya Krempeaux](http://reiver.link)
