#!/usr/bin/env python3

import os
import sys
import traceback
import datetime
import argparse

def main():
    parser = argparse.ArgumentParser(
        description="Script for generating SemVer-compatible formatted version numbers from an input date",
        formatter_class=argparse.ArgumentDefaultsHelpFormatter,
    )

    parser.add_argument(
        "datetime_string",
        help="The date string in ISO 8601 format to parse. If it's unset, the value from the environment variable CI_PIPELINE_CREATED_AT will be used.",
        type=str,
        nargs="?"
    )

    parser.add_argument(
        "--use-zero-padding",
        help="Whether to pad numbers with zeroes",
        action="store_true",
        required=False,
        default=False,
    )
    
    args = parser.parse_args()
    
    # for use with GitLab CI/CD
    # see https://docs.gitlab.com/ee/ci/variables/predefined_variables.html
    dt_str = os.getenv("CI_PIPELINE_CREATED_AT")
    
    # if it's not defined or if an argument is given, use that instead
    if (args.datetime_string):
        dt_str = args.datetime_string
        
    if not dt_str:
        print("Arg can't be empty if CI_PIPELINE_CREATED_AT is not set!")
        sys.exit(2)
    
    dt = datetime.datetime.fromisoformat(dt_str)
    
    major = dt.year
    
    if args.use_zero_padding:
        minor = f"{dt.month:02}{dt.day:02}"
        patch = f"{dt.hour:02}{dt.minute:02}{dt.second:02}"
    else:
        minor = f"{dt.month}{dt.day:02}"
        patch = f"{dt.hour}{dt.minute:02}{dt.second:02}"
        
    full = f"{major}.{minor}.{patch}"
    
    # print so that it can be redirected to a file
    sys.stdout.write(f"DATETIME_VERSION_MAJOR={major}\n")
    sys.stdout.write(f"DATETIME_VERSION_MINOR={minor}\n")
    sys.stdout.write(f"DATETIME_VERSION_PATCH={patch}\n")
    sys.stdout.write(f"DATETIME_VERSION_FULL={full}\n")


if __name__ == "__main__":
    try:
        main()
    except Exception as e:
        traceback.print_exc()
        sys.exit(1)