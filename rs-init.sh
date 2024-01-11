#!/bin/bash
mongosh <<EOF
var config = {
    "_id": "dbrs",
    "version": 1,
    "members": [
        {
            "_id": 1,
            "host": "mongowrite:27017",
            "priority": 3
        },
        {
            "_id": 2,
            "host": "mongoread:27017",
            "priority": 2
        },
    ]
};
rs.initiate(config, { force: true });
rs.status();
EOF



