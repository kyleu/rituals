namespace auth {
  export interface Auth {
    readonly id: string;
    readonly provider: string;
    readonly email: string;
    readonly providesUsers: boolean;
  }

  export interface Provider {
    readonly key: string,
    readonly title: string
  }

  const github: Provider = {key: "github", title: "GitHub"};
  const google: Provider = {key: "google", title: "Google"};
  const slack: Provider = {key: "slack", title: "Slack"};
  const amazon: Provider = {key: "amazon", title: "Amazon"};
  const microsoft: Provider = {key: "microsoft", title: "Microsoft"};

  export const allProviders = [github, google, slack, amazon, microsoft];

  export interface Email {
    readonly matched: boolean,
    readonly domain: string
  }
}

