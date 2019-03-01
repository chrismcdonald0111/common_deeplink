package service

import "testing"

func TestIsValidPortNum(t *testing.T) {
  aboveMaxPortNum := 75535
  aboveMaxFlag    := isValidPortNum(aboveMaxPortNum)
  if aboveMaxFlag != false {
     t.Errorf("Return value was incorrect, got: %t, want: %t.", aboveMaxFlag, false)
  }

  belowMinPortNum := 999
  belowMinFlag    := isValidPortNum(belowMinPortNum)
  if belowMinFlag != false {
     t.Errorf("Return value was incorrect, got: %t, want: %t.", belowMinFlag, false)
  }

  validPortNum    := 2000
  validPortFlag   := isValidPortNum(validPortNum)
  if validPortFlag != true {
     t.Errorf("Return value was incorrect, got: %t, want: %t.", validPortFlag, true)
  }
}

func TestGetPortNum(t *testing.T) {
  defaultPortNum := 8080
  invalidPortNum := 99999
  invalidPortVal := getPortNum(invalidPortNum)
  if invalidPortVal != defaultPortNum {
     t.Errorf("Return value was incorrect, got: %d, want: %d.", invalidPortVal, defaultPortNum)
  }

  validPortNum := 1111
  validPortVal := getPortNum(validPortNum)
  if validPortVal != validPortNum {
     t.Errorf("Return value was incorrect, got: %d, want: %d.", validPortVal, validPortNum)
  }
}

func TestGetTimeout(t *testing.T) {
  defaultTimeout := Timeout{
    ReadTimeoutSec:  5,
    WriteTimeoutSec: 10,
    IdleTimeoutSec:  60,
  }
  emptyTimeout := Timeout{}
  emptyTimeoutVal := getTimeout(emptyTimeout)
  if emptyTimeoutVal != defaultTimeout {
     t.Errorf("Return value was incorrect, got: %+v, want: %+v.", emptyTimeoutVal, defaultTimeout)
  }

  validTimeout := Timeout{
    ReadTimeoutSec:  5,
    WriteTimeoutSec: 15,
    IdleTimeoutSec:  30,
  }
  validTimeoutVal := getTimeout(validTimeout)
  if validTimeoutVal != validTimeout {
     t.Errorf("Return value was incorrect, got: %+v, want: %+v.", validTimeoutVal, validTimeout)
  }

  readOutOfRangeTimeout := Timeout{
    ReadTimeoutSec:  30,
    WriteTimeoutSec: 30,
    IdleTimeoutSec:  30,
  }
  expectedReadOutOfRangeTimeout := Timeout{
    ReadTimeoutSec:  5,
    WriteTimeoutSec: 30,
    IdleTimeoutSec:  30,
  }
  readOutOfRangeTimeoutVal := getTimeout(readOutOfRangeTimeout)
  if readOutOfRangeTimeoutVal != expectedReadOutOfRangeTimeout {
     t.Errorf("Return value was incorrect, got: %+v, want: %+v.", readOutOfRangeTimeoutVal, expectedReadOutOfRangeTimeout)
  }

  writeOutOfRangeTimeout := Timeout{
    ReadTimeoutSec:  5,
    WriteTimeoutSec: 50,
    IdleTimeoutSec:  30,
  }
  expectedWriteOutOfRangeTimeout := Timeout{
    ReadTimeoutSec:  5,
    WriteTimeoutSec: 10,
    IdleTimeoutSec:  30,
  }
  writeOutOfRangeTimeoutVal := getTimeout(writeOutOfRangeTimeout)
  if writeOutOfRangeTimeoutVal != expectedWriteOutOfRangeTimeout {
     t.Errorf("Return value was incorrect, got: %+v, want: %+v.", writeOutOfRangeTimeoutVal, expectedWriteOutOfRangeTimeout)
  }

  idleOutOfRangeTimeout := Timeout{
    ReadTimeoutSec:  5,
    WriteTimeoutSec: 10,
    IdleTimeoutSec:  200,
  }
  expectedIdleOutOfRangeTimeout := Timeout{
    ReadTimeoutSec:  5,
    WriteTimeoutSec: 10,
    IdleTimeoutSec:  60,
  }
  idleOutOfRangeTimeoutVal := getTimeout(idleOutOfRangeTimeout)
  if idleOutOfRangeTimeoutVal != expectedIdleOutOfRangeTimeout {
     t.Errorf("Return value was incorrect, got: %+v, want: %+v.", idleOutOfRangeTimeoutVal, expectedIdleOutOfRangeTimeout)
  }
}
