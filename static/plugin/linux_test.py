import json
import argparse
import requests

headers = {
    "User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
    "Origin": "https://leetcode.cn",
    "Host": "leetcode.cn",
    "Content-Type": "application/json"
}


def get_usr_code_list(name):
    url = "https://leetcode.cn/graphql/noj-go/"

    param_data = {
        "query": "query recentAcSubmissions($userSlug: String!) {recentACSubmissions(userSlug: $userSlug) {"
                 "submissionId submitTime question {title translatedTitle titleSlug questionFrontendId}}}",
        "variables": {
            "userSlug": name
        },
        "operationName": "recentAcSubmissions"
    }

    content = requests.post(url, headers=headers, data=json.dumps(param_data))

    print(json.dumps(content.json(), indent=4))


def run(name):
    get_usr_code_list(name)


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument('--name', type=str, default = "newtsun")
    args = parser.parse_args()
    run(args.name)
