# Emailcheck

Emailcheck is a go cli application that will authenticate with Gmail and check for emails based on a filter. If there are emails matching the filter than a new email will be sent to a notification address alerting it of the new emails

## Use Case
I'm using this to watch my inbox and notify me when I get new emails from specific recruiters or employers via the SMS/MMS gateway from my mobile phone provider.

## Pre-requesites
* [Go](https://golang.org/), latest version recommended.
* A Google Cloud Platform project with the API enabled. To create a project and enable an API, refer to [Create a project and enable the API](https://developers.google.com/workspace/guides/create-project)
* Authorization credentials for a desktop application. To learn how to create credentials for a desktop application, refer to [Create credentials](https://developers.google.com/workspace/guides/create-credentials).
* A Google account with Gmail enabled.

## Setup
* Add json files to the "configs/checks" directory (see example.json.sample)
* Download credentials.json file from Google Cloud Platform project and store in the configs directory

## Run
* Execute `./bin/emailcheck`
  * On first run, it will output a link. After following this and authenticating, copy the code from the URL and paste it in your terminal and press ENTER
