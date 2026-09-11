node1 :) SELECT uuid FROM system.databases WHERE database='r';
node4 :) CREATE DATABASE r UUID '<uuid from previous query>' ENGINE=Replicated('some/path/{uuid}','other_shard','r2');
