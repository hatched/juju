#!/usr/bin/env python
from __future__ import print_function

from argparse import ArgumentParser
import json
import os

from jenkins import Jenkins
from jujuci import (
    add_credential_args,
    get_credentials,
    )
from utility import (
    find_candidates,
    get_auth_token,
    get_candidates_path,
    )


def get_args(argv=None):
    parser = ArgumentParser()
    parser.add_argument(
        'root_dir', help='Directory containing releases and candidates dir')
    parser.add_argument(
        '--all', action='store_true', default=False,
        help='Schedule all candidates for client-server testing.')
    add_credential_args(parser)
    args = parser.parse_args(argv)
    return args, get_credentials(args)


def get_releases(root):
    release_path = os.path.join(root, 'old-juju')
    for entry in os.listdir(release_path):
        if not os.path.isdir(os.path.join(release_path, entry)):
            continue
        yield entry


def get_candidate_version(candidate_path):
    with open(os.path.join(candidate_path, 'buildvars.json')) as fp:
        return json.load(fp)['version']


def calculate_jobs(root, schedule_all=False):
    releases = list(get_releases(root))
    candidates_path = get_candidates_path(root)
    for candidate_path in find_candidates(root, schedule_all):
        parent, candidate = os.path.split(candidate_path)
        if parent != candidates_path:
            raise ValueError('Wrong path')
        candidate_version = get_candidate_version(candidate_path)
        for release in releases:
            if is_osx_client(candidate) is not is_osx_client(release):
                continue
            client_os = 'ubuntu'
            if is_osx_client(release):
                client_os = 'osx'
            yield {
                'old_version': release,  # Client
                'candidate': candidate_version,  # Server
                'new_to_old': 'true',
                'candidate_path': candidate,
                'client_os': client_os,
            }
            yield {
                'old_version': release,  # Server
                'candidate': candidate_version,  # Client
                'new_to_old': 'false',
                'candidate_path': candidate,
                'client_os': client_os,
            }


def is_osx_client(path):
    return path.endswith('-osx')


def build_jobs(credentials, root, jobs):
    jenkins = Jenkins('http://localhost:8080', *credentials)
    token = get_auth_token(root, 'compatibility-control')
    for job in jobs:
        jenkins.build_job('compatibility-control', job, token=token)


def main():
    args, credentials = get_args()
    build_jobs(
        credentials, args.root_dir, calculate_jobs(args.root_dir, args.all))


if __name__ == '__main__':
    main()
