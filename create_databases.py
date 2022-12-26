#! /usr/bin/env python3
import os
import sys
import time
import socket
import subprocess

timeout = 90
host = 'localhost'
port = 54321

user = os.getenv('POSTGRES_USER')

def run(command):
    print(command)
    proc = subprocess.Popen([command], stdout=subprocess.PIPE, stderr=subprocess.PIPE, shell=True)
    return proc.communicate()


def create_db(db):
    out, err = run('createdb -U %s %s' % (user, db))
    if err and b'already exists' not in err:
        return out, err
    return run('psql -U %s -d %s -c "CREATE EXTENSION IF NOT EXISTS pg_trgm;"' % (user, db))


def wait():
    start = time.time()
    while True:
        print('Waiting postgres on 5432...')
        try:
            sock = socket.socket()
            sock.connect((host, port))
            sock.close()
            break
        except socket.error as e:
            if time.time() - start >= timeout:
                print('Timeout exceeded')
                sys.exit(1)

            time.sleep(3)


databases = os.getenv('POSTGRES_DBS').split(',')
wait()
out, err = create_db(db)
if b'the database system is starting up' in err:
    time.sleep(10)
    break
if err and b'already exists' not in err:
    print(err)