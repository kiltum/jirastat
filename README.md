# Jira Statistic data generator

Common JIRA report "Created vs Resolved Issues Report" not feet my needs.

I need "Created vs {STATUS} Issues Report" instead. But JIRA does not give me such instruments

So here is very simple Go program to do this job. 

_Compiled jirastat in repo - for OS X_

## Usage

Set environment variables or command line options and run. You will get something like this:

```
JS_PASS=pass JS_CUMULATIVE=no ./jirastat --js_host https://atlassin.net --js_user admin

Date	Created	Updated
-----------------------
2018-12-25	5	31
2018-12-26	9	12
2018-12-27	11	8
2018-12-28	9	5
2018-12-29	4	5
2019-01-09	13	11
2019-01-10	11	9
2019-01-11	15	9
2019-01-12	2	6
2019-01-13	1	1
2019-01-14	21	7
```

Now you can simple copy and paste to excel (or other) and make any graph you need. For your purpose i already create simple book.xls

## Configure


|Name|Default|Description|
|----|-------|-----------|
|JS_HOST|none|URL of jira server|
|JS_USER|none|Username|
|JS_PASS|none|Password|
|JS_PROJECT|IT|Short key of project in JIRA|
|JS_STATUS|DONE|Which status should count for updated?|
|JS_DAYS|30|How many days script should look in past|
|JS_CUMILATIVE|yes|Cumulative output in task count|
|JS_VERB|no|Put some debug information

Also all this config variables can be storen in jirastat.{json,yaml,toml} file

## Example picture from Excel

![alt text](https://raw.githubusercontent.com/kiltum/jirastat/master/pic.png)