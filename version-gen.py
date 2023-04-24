#!/usr/bin/env python3

import os
import sys
import traceback
import datetime


def main():
    # for use with GitLab CI/CD
    # see https://docs.gitlab.com/ee/ci/variables/predefined_variables.html
    dt_str = os.getenv("CI_PIPELINE_CREATED_AT")
    
    # if it's not defined or if an argument is given, use that instead
    if (len(sys.argv) == 2):
        dt_str = sys.argv[1]
        
    if not dt_str:
        print("Arg can't be empty if CI_PIPELINE_CREATED_AT is not set!")
        sys.exit(2)
    
    dt = datetime.datetime.fromisoformat(dt_str)
    
    major = dt.year
    minor = f"{dt.month}{dt.day:02}"
    patch = f"{dt.hour}{dt.minute:02}{dt.second:02}"
    
    # print so that it can be redirected to a file
    print(f"DATETIME_VERSION_MAJOR={major}")
    print(f"DATETIME_VERSION_MINOR={minor}")
    print(f"DATETIME_VERSION_PATCH={patch}")
    print(f"DATETIME_VERSION_FULL={major}.{minor}.{patch}", flush=True)


if __name__ == "__main__":
    try:
        main()
    except Exception as e:
        traceback.print_exc()
        sys.exit(1)