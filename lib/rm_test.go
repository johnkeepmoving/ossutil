package lib 

import (
    "fmt"

    . "gopkg.in/check.v1"
)

func (s *OssutilCommandSuite) rawRemove(args []string, recursive, force, bucket bool) (bool, error) {
    command := "rm"
    str := ""
    options := OptionMapType{
        "endpoint": &str,
        "accessKeyID": &str,
        "accessKeySecret": &str,
        "stsToken": &str,
        "configFile": &configFile,
        "recursive": &recursive,
        "force": &force,
        "bucket": &bucket,
    }
    showElapse, err := cm.RunCommand(command, args, options)
    return showElapse, err
}

func (s *OssutilCommandSuite) TestRemoveObject(c *C) {
    bucket := bucketNamePrefix + "rm"
    s.putBucket(bucket, c)

    // put object
    object := "test_object"
    s.putObject(bucket, object, uploadFileName, c)

    // list object
    objects := s.listObjects(bucket, "", false, false, c)
    c.Assert(len(objects), Equals, 1)
    c.Assert(objects[0], Equals, object)

    // remove object
    s.removeObjects(bucket, object, false, true, c)

    // list object
    objects = s.listObjects(bucket, "", false, false, c)
    c.Assert(len(objects), Equals, 0)
}

func (s *OssutilCommandSuite) TestRemoveObjects(c *C) {
    bucket := bucketNamePrefix + "rm"
    s.putBucket(bucket, c)

    // put object
    num := 15
    objectNames := []string{}
    for i := 0; i < num; i++ {
        object := fmt.Sprintf("^$%d$^工具", i) 
        s.putObject(bucket, object, uploadFileName, c) 
        objectNames = append(objectNames, object)
    }

    // list object
    objects := s.listObjects(bucket, "", false, false, c)
    c.Assert(len(objects), Equals, num)

    // "rm oss://bucket/ -r"
    // remove object
    s.removeObjects(bucket, "", true, false, c)

    objects = s.listObjects(bucket, "", false, false, c)
    c.Assert(len(objects), Equals, num)

    // "rm oss://bucket/prefix -r -f"
    // remove object
    s.removeObjects(bucket, "^$", true, true, c)

    // list object
    objects = s.listObjects(bucket, "", false, false, c)
    c.Assert(len(objects), Equals, 0)
}

func (s *OssutilCommandSuite) TestRemoveBucket(c *C) {
    bucket := bucketNamePrefix + "rm"
    s.putBucket(bucket, c)

    // put object
    num := 15
    objectNames := []string{}
    for i := 0; i < num; i++ {
        object := fmt.Sprintf("^$%d$^工具", i) 
        s.putObject(bucket, object, uploadFileName, c) 
        objectNames = append(objectNames, object)
    }
    
    // list buckets
    buckets := s.listBuckets(false, c)
    c.Assert(FindPos(bucket, buckets) != -1, Equals, true)

    // rm bucket
    s.removeBucket(bucket, true, c)

    // list buckets
    buckets = s.listBuckets(false, c)
    c.Assert(FindPos(bucket, buckets) == -1, Equals, true)
}

func (s *OssutilCommandSuite) TestRemoveEmptyBucket(c *C) {
    bucket := bucketNamePrefix + "rm"
    s.putBucket(bucket, c)

    // list buckets
    buckets := s.listBuckets(false, c)
    c.Assert(FindPos(bucket, buckets) != -1, Equals, true)

    // rm bucket
    s.removeBucket(bucket, false, c)

    // list buckets
    buckets = s.listBuckets(false, c)
    c.Assert(FindPos(bucket, buckets) == -1, Equals, true)
}

func (s *OssutilCommandSuite) TestRemoveNonEmptyBucket(c *C) {
    bucket := bucketNamePrefix + "rm"
    s.putBucket(bucket, c)

    // put object
    object := "test_object"
    s.putObject(bucket, object, uploadFileName, c)

    command := "rm"
    args := []string{CloudURLToString(bucket, "")}
    str := ""
    ok := true
    options := OptionMapType{
        "endpoint": &str,
        "accessKeyID": &str,
        "accessKeySecret": &str,
        "stsToken": &str,
        "configFile": &configFile,
        "bucket": &ok,
        "force": &ok,
    }
    _, err := cm.RunCommand(command, args, options)
    c.Assert(err, NotNil)

    // list object
    objects := s.listObjects(bucket, "", false, false, c)
    c.Assert(len(objects), Equals, 1)
    c.Assert(objects[0], Equals, object)
}

func (s *OssutilCommandSuite) TestRemoveObjectBucketOption(c *C) {
    bucket := bucketNamePrefix + "rm"
    s.putBucket(bucket, c)

    object := "test_object"
    command := "rm"
    args := []string{CloudURLToString(bucket, object)}
    str := ""
    ok := true
    options := OptionMapType{
        "endpoint": &str,
        "accessKeyID": &str,
        "accessKeySecret": &str,
        "stsToken": &str,
        "configFile": &configFile,
        "bucket": &ok,
        "force": &ok,
    }
    _, err := cm.RunCommand(command, args, options)
    c.Assert(err, NotNil)

    // list buckets
    buckets := s.listBuckets(false, c)
    c.Assert(FindPos(bucket, buckets) != -1, Equals, true)
}

func (s *OssutilCommandSuite) TestErrRemove(c *C) {
    bucket := bucketNamePrefix + "rm"
    s.putBucket(bucket, c)

    showElapse, err := s.rawRemove([]string{"oss://"}, false, true, true)
    c.Assert(err, NotNil)
    c.Assert(showElapse, Equals, false)

    showElapse, err = s.rawRemove([]string{"./"}, false, true, true)
    c.Assert(err, NotNil)
    c.Assert(showElapse, Equals, false)

    showElapse, err = s.rawRemove([]string{CloudURLToString(bucket, "")}, false, true, false)
    c.Assert(err, NotNil)
    c.Assert(showElapse, Equals, false)

    object := "/object"
    showElapse, err = s.rawRemove([]string{CloudURLToString(bucket, object)}, false, true, false)
    c.Assert(err, NotNil)
    c.Assert(showElapse, Equals, false)

    // remove bucket without force
    showElapse, err = s.rawRemove([]string{CloudURLToString(bucket, "")}, false, false, true)
    c.Assert(err, IsNil)
    c.Assert(showElapse, Equals, true)

    bucketStat := s.getStat(bucket, "", c)
    c.Assert(bucketStat[StatName], Equals, bucket)

    // batch delete not exist objects
    showElapse, err = s.rawRemove([]string{CloudURLToString(bucket, object)}, true, true, false)
    c.Assert(err, NotNil)
    c.Assert(showElapse, Equals, false)

    // clear not exist bucket
    bucketName := bucketNamePrefix + "rmnotexist"
    showElapse, err = s.rawRemove([]string{CloudURLToString(bucketName, "")}, true, true, false)
    c.Assert(err, NotNil)
    c.Assert(showElapse, Equals, false)
}

func (s *OssutilCommandSuite) TestErrDeleteObject(c *C) {
    bucketName := bucketNamePrefix + "rm"

    bucket, err := removeCommand.command.ossBucket(bucketName)
    c.Assert(err, IsNil)

    object := "object"
    err = removeCommand.ossDeleteObjectRetry(bucket, object)
    c.Assert(err, NotNil)

    _, err = removeCommand.ossBatchDeleteObjectsRetry(bucket, []string{object})
    c.Assert(err, NotNil)
}
