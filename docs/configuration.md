
# How to configure Umotd

Umotd has a default built-in look, but it's actually made to be highly customizable.
If you're managing a custom system, it might interest you.

## Where to put your config

You can create a custom config file at `/etc/umotd/config.json` (system-wide) or `~/.config/umotd/config.json` (user-specific).

## Translations ?!

Umotd supports translations for specific tips, command descriptions and link names.
They have specific names / codes that are used to get translated strings.

Any other option not listed won't be translated.

## Breaking down the configuration file

Here's a breakdown of the config file options - there's also the example folder if you want to see concrete use cases.

## Commands

This option allows you to define a list of commands to display in the MOTD.

Here are the unique codes you can use to get translated strings for command descriptions : 

- `cmd_list`: "List of available commands"
- `cli_pkg`: "Manage command line packages"
- `term_bling`: "Enable terminal bling"
- `motd_toggle`: "Toggle this banner on/off" (there are no built-in commands for this)
- `sys_info`: "View system info"
- `man_upd`: "Manually update the system"

```json
{
  "commands": [
    {
      "cmd": "ujust aurora-cli",
      "desc": "term_bling"
    },
    {
      "cmd": "fastfetch",
      "desc": "sys_info"
    },
    {
      "cmd": "brew help",
      "desc": "cli_pkg"
    },
    {
      "cmd": "cowsay",
      "desc": "Display a cow saying something"
    }
  ]
}
```

## Links

This option allows to add custom links to the MOTD.

There are unique names you can use to get a translated name for the link :

- `website` : "Website"
- `issues` : "Report an issue"
- `docs` : "Documentation"
- `discuss` : "Discuss"
- `discord` : "Discord"
- `matrix` : "Matrix"
- `bluesky` : "Bluesky"
- `mastodon` : "Mastodon"
- `donate` : "Donate"

```json
{
  "links": [
    {
      "name": "issues",
      "url": "https://issues.bazzite.gg/"
    },
    {
      "name": "docs",
      "url": "https://docs.bazzite.gg/"
    },
    {
      "name": "discord",
      "url": "https://discord.gg/bazzite"
    },
    {
      "name": "bluesky",
      "url": "https://bluesky.bazzite.gg/"
    },
    {
      "name": "discuss",
      "url": "https://github.com/ublue-os/aurora/discussions"
    },
    {
      "name": "Custom Link (it won't get translated)",
      "url": "https://www.innersloth.com/games/among-us/"
    }
  ]
}
```

## Prefix and Suffix

These options allow to customize the prefix and suffix of the welcome message.

```json
{
  "prefix": "> ",
  "suffix": " !"
}
```

Example:

```
`> Welcome to UBlue !`
```

## Tips

The `tips` option allows to add custom tips to the MOTD.

```json
{
  "tips": [
    "This is a custom tip by yours truly ! :D",
    "This is another custom tip (they won't get translated)"
  ]
}
```

### Tips Presets

The `tips-presets` option also allows to use the predefined and translated tips included in umotd.

Tips presets can also be combined to show new tips alongside the previous ones. For example : `ublue` + `dev` displays exclusive dev tips for the ublue systems alongside the normal `ublue` and `dev` tips.

Currently, there are the following presets available:

- `aurora`
- `bazzite`
- `deck`
- `bluefin`
- `default`
- `gnome`
- `kde`
- `ublue`
- `dev`

```json
{
  "tips-presets": [
    "gnome",
    "bluefin",
    "ublue",
    "dev"
  ]
}
```

## Use Accent Color

This option allows umotd to use the accent color of the system.

> Note: It's only available for the GNOME desktop as it relies on `gsettings`.

```json
{
  "use-accent-color": true
}
```

## Info File

This option allows to redirect the info file path to a custom location.

> Note: This option is specifically tailored for bootc / ublue images, because it uses the image info file of their system to display information about the system image the user is currently running.


```json
{
  "info-file": "/usr/share/ublue-os/image-info.json"
}
```
