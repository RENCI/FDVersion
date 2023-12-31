# FDVersion


This tool is purposed to collect and aggregate information about files and directories. 

Example, how to run:

`FDVersion.exe -path="C:\repo\FDVersion\tests\data\case07_hashes" -output="./res.xml"`

_path - path to directory for processing_

_output - path for output file_

Example of an output file:
```
 <NodeXML id="C:\repo\FDVersion\tests\data\case07_hashes">
    <Name>case07_hashes</Name>
    <Path>C:\repo\FDVersion\tests\data\case07_hashes</Path>
    <IsDir>true</IsDir>
    <IsFile>false</IsFile>
    <FileSize>0</FileSize>
    <Hash>63171381bbf6710fbde74ba3d53ff6df03b2bb853d565d0096960c5fb31fb493</Hash>
    <Nodes>
        <NodeXML id="C:\repo\FDVersion\tests\data\case07_hashes\file1.txt">
            <Name>file1.txt</Name>
            <Path>C:\repo\FDVersion\tests\data\case07_hashes\file1.txt</Path>
            <IsDir>false</IsDir>
            <IsFile>true</IsFile>
            <FileSize>28</FileSize>
            <Hash>47f8cf7a375b87e5b1ad96a2b82b594b064b97900196820f65741cd2ebe9caae</Hash>
            <Nodes></Nodes>
        </NodeXML>
        <NodeXML id="C:\repo\FDVersion\tests\data\case07_hashes\dir1">
            <Name>dir1</Name>
            <Path>C:\repo\FDVersion\tests\data\case07_hashes\dir1</Path>
            <IsDir>true</IsDir>
            <IsFile>false</IsFile>
            <FileSize>0</FileSize>
            <Hash>85f3e2903210f99e678a3c2cf48be0be2b9a4c0c08403f00535f21c1301743c8</Hash>
            <Nodes>
                <NodeXML id="C:\repo\FDVersion\tests\data\case07_hashes\dir1\file2.txt">
                    <Name>file2.txt</Name>
                    <Path>C:\repo\FDVersion\tests\data\case07_hashes\dir1\file2.txt</Path>
                    <IsDir>false</IsDir>
                    <IsFile>true</IsFile>
                    <FileSize>4</FileSize>
                    <Hash>3b9c358f36f0a31b6ad3e14f309c7cf198ac9246e8316f9ce543d5b19ac02b80</Hash>
                    <Nodes></Nodes>
                </NodeXML>
                <NodeXML id="C:\repo\FDVersion\tests\data\case07_hashes\dir1\file3.txt">
                    <Name>file3.txt</Name>
                    <Path>C:\repo\FDVersion\tests\data\case07_hashes\dir1\file3.txt</Path>
                    <IsDir>false</IsDir>
                    <IsFile>true</IsFile>
                    <FileSize>5</FileSize>
                    <Hash>6f3fef6dc51c7996a74992b70d0c35f328ed909a5e07646cf0bab3383c95bb02</Hash>
                    <Nodes></Nodes>
                </NodeXML>
            </Nodes>
        </NodeXML>
    </Nodes>
</NodeXML>
```

## Hash

For each file and directory we calculate Hash value. 

SHA256 algorithm is used.

File content is hashed for a file.

Directory hash is a hash value of hashes of all child elements, so change in any of child elements will result in updated value of the parent(s) directories.

