(module
  (type (;0;) (func (param i32 i32 i32 i32)))
  (type (;1;) (func (param i32 i32) (result i32)))
  (type (;2;) (func (param i32 i32 i32)))
  (type (;3;) (func (param i32 i32)))
  (type (;4;) (func (param i64) (result i32)))
  (type (;5;) (func (result i32)))
  (type (;6;) (func (param i32 i32 i32) (result i32)))
  (type (;7;) (func (result i64)))
  (type (;8;) (func (param i32) (result i32)))
  (type (;9;) (func (param i32 i32 i64 i32 i32) (result i32)))
  (type (;10;) (func (param i32)))
  (type (;11;) (func (param i32) (result i64)))
  (type (;12;) (func (param i32 i32 i32 i32) (result i32)))
  (type (;13;) (func (param i64 i32 i32 i32 i32 i32) (result i32)))
  (type (;14;) (func))
  (type (;15;) (func (param i64)))
  (type (;16;) (func (param i32 i32 i32 i32 i32)))
  (type (;17;) (func (param i32 i32 i32 i32 i32 i32 i32) (result i64)))
  (type (;18;) (func (param i64 i32 i32) (result i32)))
  (type (;19;) (func (param i64 i32 i32 i32 i32) (result i32)))
  (type (;20;) (func (param i32 i32 i64 i32)))
  (type (;21;) (func (param i32 i32 i64 i32 i32 i32)))
  (type (;22;) (func (param i32 i32 i64)))
  (type (;23;) (func (param i32 i64)))
  (type (;24;) (func (param i32 i32) (result i64)))
  (type (;25;) (func (param i32 i64) (result i32)))
  (type (;26;) (func (param i32 i32 i32 i32 i32) (result i32)))
  (import "env" "bigIntAdd" (func (;0;) (type 2)))
  (import "env" "signalError" (func (;1;) (type 3)))
  (import "env" "bigIntNew" (func (;2;) (type 4)))
  (import "env" "mBufferNew" (func (;3;) (type 5)))
  (import "env" "mBufferAppend" (func (;4;) (type 1)))
  (import "env" "mBufferSetBytes" (func (;5;) (type 6)))
  (import "env" "getGasLeft" (func (;6;) (type 7)))
  (import "env" "mBufferNewFromBytes" (func (;7;) (type 1)))
  (import "env" "mBufferGetLength" (func (;8;) (type 8)))
  (import "env" "managedTransferValueExecute" (func (;9;) (type 9)))
  (import "env" "managedCaller" (func (;10;) (type 10)))
  (import "env" "managedOwnerAddress" (func (;11;) (type 10)))
  (import "env" "mBufferEq" (func (;12;) (type 1)))
  (import "env" "getNumDCDTTransfers" (func (;13;) (type 5)))
  (import "env" "mBufferGetArgument" (func (;14;) (type 1)))
  (import "env" "mBufferAppendBytes" (func (;15;) (type 6)))
  (import "env" "managedSignalError" (func (;16;) (type 10)))
  (import "env" "bigIntGetUnsignedArgument" (func (;17;) (type 3)))
  (import "env" "smallIntGetUnsignedArgument" (func (;18;) (type 11)))
  (import "env" "getNumArguments" (func (;19;) (type 5)))
  (import "env" "managedGetOriginalTxHash" (func (;20;) (type 10)))
  (import "env" "mBufferToBigIntUnsigned" (func (;21;) (type 1)))
  (import "env" "bigIntSign" (func (;22;) (type 8)))
  (import "env" "mBufferFromBigIntUnsigned" (func (;23;) (type 1)))
  (import "env" "mBufferCopyByteSlice" (func (;24;) (type 12)))
  (import "env" "mBufferStorageLoad" (func (;25;) (type 1)))
  (import "env" "mBufferStorageStore" (func (;26;) (type 1)))
  (import "env" "bigIntCmp" (func (;27;) (type 1)))
  (import "env" "managedSCAddress" (func (;28;) (type 10)))
  (import "env" "managedAsyncCall" (func (;29;) (type 0)))
  (import "env" "managedMultiTransferDCDTNFTExecute" (func (;30;) (type 9)))
  (import "env" "managedExecuteOnDestContext" (func (;31;) (type 13)))
  (import "env" "mBufferSetRandom" (func (;32;) (type 1)))
  (import "env" "mBufferGetBytes" (func (;33;) (type 1)))
  (import "env" "checkNoPayment" (func (;34;) (type 14)))
  (import "env" "bigIntGetDCDTCallValue" (func (;35;) (type 10)))
  (import "env" "getDCDTTokenName" (func (;36;) (type 8)))
  (import "env" "getBlockTimestamp" (func (;37;) (type 7)))
  (import "env" "bigIntMul" (func (;38;) (type 2)))
  (import "env" "bigIntGetExternalBalance" (func (;39;) (type 3)))
  (import "env" "bigIntFinishUnsigned" (func (;40;) (type 10)))
  (import "env" "smallIntFinishUnsigned" (func (;41;) (type 15)))
  (import "env" "bigIntGetCallValue" (func (;42;) (type 10)))
  (import "env" "mBufferGetByteSlice" (func (;43;) (type 12)))
  (func (;44;) (type 8) (param i32) (result i32)
    (local i32)
    call 45
    local.tee 1
    local.get 1
    local.get 0
    call 0
    local.get 1)
  (func (;45;) (type 5) (result i32)
    i64.const 0
    call 2)
  (func (;46;) (type 3) (param i32 i32)
    local.get 0
    local.get 1
    call 1
    unreachable)
  (func (;47;) (type 16) (param i32 i32 i32 i32 i32)
    block  ;; label = @1
      block  ;; label = @2
        local.get 2
        local.get 1
        i32.lt_u
        br_if 0 (;@2;)
        local.get 2
        local.get 4
        i32.le_u
        br_if 1 (;@1;)
        local.get 2
        local.get 4
        call 48
        unreachable
      end
      local.get 1
      local.get 2
      call 49
      unreachable
    end
    local.get 0
    local.get 2
    local.get 1
    i32.sub
    i32.store offset=4
    local.get 0
    local.get 3
    local.get 1
    i32.add
    i32.store)
  (func (;48;) (type 3) (param i32 i32)
    local.get 0
    local.get 1
    call 184
    unreachable)
  (func (;49;) (type 3) (param i32 i32)
    local.get 0
    local.get 1
    call 187
    unreachable)
  (func (;50;) (type 5) (result i32)
    i64.const 1
    call 2)
  (func (;51;) (type 8) (param i32) (result i32)
    (local i32)
    call 3
    local.tee 1
    local.get 0
    call 4
    drop
    local.get 1)
  (func (;52;) (type 3) (param i32 i32)
    local.get 0
    local.get 1
    call 46
    unreachable)
  (func (;53;) (type 17) (param i32 i32 i32 i32 i32 i32 i32) (result i64)
    (local i32 i32 i64)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 7
    global.set 0
    call 3
    local.tee 8
    local.get 0
    call 54
    local.get 8
    local.get 1
    call 55
    local.get 8
    local.get 2
    call 56
    local.get 8
    local.get 3
    call 55
    local.get 8
    local.get 4
    call 56
    local.get 8
    local.get 5
    call 56
    block  ;; label = @1
      block  ;; label = @2
        local.get 6
        i32.load
        local.tee 5
        call 57
        i32.eqz
        br_if 0 (;@2;)
        local.get 7
        local.get 5
        call 57
        i32.store offset=24
        local.get 7
        i32.const 0
        i32.store offset=20
        local.get 7
        local.get 6
        i32.store offset=16
        loop  ;; label = @3
          local.get 7
          i32.const 8
          i32.add
          local.get 7
          i32.const 16
          i32.add
          call 58
          local.get 7
          i32.load offset=8
          i32.eqz
          br_if 2 (;@1;)
          local.get 8
          local.get 7
          i32.load offset=12
          call 59
          br 0 (;@3;)
        end
      end
      call 3
      local.tee 6
      i32.const 1049856
      i32.const 0
      call 5
      drop
      local.get 8
      local.get 6
      call 60
    end
    call 6
    i32.const 1048639
    i32.const 13
    call 7
    local.get 8
    call 61
    local.set 8
    local.get 7
    i32.const 0
    i32.store offset=16
    i64.const 0
    local.set 9
    block  ;; label = @1
      local.get 8
      i32.const 0
      local.get 7
      i32.const 16
      i32.add
      i32.const 4
      call 62
      br_if 0 (;@1;)
      local.get 7
      i32.load offset=16
      local.tee 8
      i32.const 24
      i32.shl
      local.get 8
      i32.const 8
      i32.shl
      i32.const 16711680
      i32.and
      i32.or
      local.get 8
      i32.const 8
      i32.shr_u
      i32.const 65280
      i32.and
      local.get 8
      i32.const 24
      i32.shr_u
      i32.or
      i32.or
      local.tee 6
      call 8
      local.tee 8
      i32.const 8
      i32.gt_u
      br_if 0 (;@1;)
      i64.const 0
      local.set 9
      local.get 7
      i64.const 0
      i64.store offset=16
      local.get 6
      i32.const 0
      local.get 7
      i32.const 16
      i32.add
      local.get 8
      i32.sub
      i32.const 8
      i32.add
      local.get 8
      call 62
      br_if 0 (;@1;)
      local.get 7
      i64.load offset=16
      local.tee 9
      i64.const 56
      i64.shl
      local.get 9
      i64.const 40
      i64.shl
      i64.const 71776119061217280
      i64.and
      i64.or
      local.get 9
      i64.const 24
      i64.shl
      i64.const 280375465082880
      i64.and
      local.get 9
      i64.const 8
      i64.shl
      i64.const 1095216660480
      i64.and
      i64.or
      i64.or
      local.get 9
      i64.const 8
      i64.shr_u
      i64.const 4278190080
      i64.and
      local.get 9
      i64.const 24
      i64.shr_u
      i64.const 16711680
      i64.and
      i64.or
      local.get 9
      i64.const 40
      i64.shr_u
      i64.const 65280
      i64.and
      local.get 9
      i64.const 56
      i64.shr_u
      i64.or
      i64.or
      i64.or
      local.set 9
    end
    local.get 7
    i32.const 32
    i32.add
    global.set 0
    local.get 9)
  (func (;54;) (type 3) (param i32 i32)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    call 3
    i32.store offset=12
    local.get 1
    local.get 2
    i32.const 12
    i32.add
    call 89
    local.get 0
    local.get 2
    i32.load offset=12
    call 60
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;55;) (type 3) (param i32 i32)
    call 3
    drop
    local.get 0
    local.get 1
    call 112
    call 60)
  (func (;56;) (type 3) (param i32 i32)
    call 3
    drop
    local.get 0
    local.get 1
    call 51
    call 60)
  (func (;57;) (type 8) (param i32) (result i32)
    local.get 0
    call 8)
  (func (;58;) (type 3) (param i32 i32)
    (local i32 i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    block  ;; label = @1
      block  ;; label = @2
        local.get 1
        i32.load offset=4
        local.tee 3
        i32.const 4
        i32.add
        local.tee 4
        local.get 1
        i32.load offset=8
        i32.le_u
        br_if 0 (;@2;)
        i32.const 0
        local.set 1
        br 1 (;@1;)
      end
      local.get 2
      i32.const 0
      i32.store offset=12
      local.get 1
      i32.load
      i32.load
      local.get 3
      local.get 2
      i32.const 12
      i32.add
      i32.const 4
      call 124
      drop
      local.get 2
      i32.load offset=12
      local.set 3
      local.get 1
      local.get 4
      i32.store offset=4
      local.get 3
      i32.const 24
      i32.shl
      local.get 3
      i32.const 8
      i32.shl
      i32.const 16711680
      i32.and
      i32.or
      local.get 3
      i32.const 8
      i32.shr_u
      i32.const 65280
      i32.and
      local.get 3
      i32.const 24
      i32.shr_u
      i32.or
      i32.or
      local.set 3
      i32.const 1
      local.set 1
    end
    local.get 0
    local.get 3
    i32.store offset=4
    local.get 0
    local.get 1
    i32.store
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;59;) (type 3) (param i32 i32)
    call 3
    drop
    local.get 0
    local.get 1
    call 51
    call 60)
  (func (;60;) (type 3) (param i32 i32)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    local.get 1
    i32.const 24
    i32.shl
    local.get 1
    i32.const 8
    i32.shl
    i32.const 16711680
    i32.and
    i32.or
    local.get 1
    i32.const 8
    i32.shr_u
    i32.const 65280
    i32.and
    local.get 1
    i32.const 24
    i32.shr_u
    i32.or
    i32.or
    i32.store offset=12
    local.get 0
    local.get 2
    i32.const 12
    i32.add
    i32.const 4
    call 15
    drop
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;61;) (type 18) (param i64 i32 i32) (result i32)
    local.get 0
    call 63
    call 45
    local.get 1
    local.get 2
    call 64)
  (func (;62;) (type 12) (param i32 i32 i32 i32) (result i32)
    local.get 0
    local.get 1
    local.get 3
    local.get 2
    call 43
    i32.const 0
    i32.ne)
  (func (;63;) (type 5) (result i32)
    (local i32)
    call 3
    local.tee 0
    call 28
    local.get 0)
  (func (;64;) (type 19) (param i64 i32 i32 i32 i32) (result i32)
    (local i32)
    local.get 0
    local.get 1
    local.get 2
    local.get 3
    local.get 4
    call 3
    local.tee 5
    call 31
    drop
    local.get 5)
  (func (;65;) (type 20) (param i32 i32 i64 i32)
    (local i32 i32)
    i32.const 1049856
    i32.const 0
    call 7
    local.set 4
    call 3
    local.set 5
    block  ;; label = @1
      local.get 1
      call 66
      br_if 0 (;@1;)
      block  ;; label = @2
        local.get 2
        i64.eqz
        i32.eqz
        br_if 0 (;@2;)
        local.get 0
        local.get 1
        i64.const 0
        local.get 3
        local.get 4
        local.get 5
        call 67
        return
      end
      local.get 0
      local.get 1
      local.get 2
      local.get 3
      local.get 4
      local.get 5
      call 67
      return
    end
    local.get 0
    local.get 3
    i64.const 0
    local.get 4
    local.get 5
    call 9
    drop)
  (func (;66;) (type 8) (param i32) (result i32)
    local.get 0
    call 57
    i32.eqz)
  (func (;67;) (type 21) (param i32 i32 i64 i32 i32 i32)
    (local i32 i32)
    global.get 0
    i32.const 48
    i32.sub
    local.tee 6
    global.set 0
    call 3
    local.set 7
    local.get 6
    local.get 1
    call 51
    local.get 2
    local.get 3
    call 44
    call 91
    local.get 6
    i32.const 32
    i32.add
    i64.const 0
    i64.store
    local.get 6
    i64.const 0
    i64.store offset=24
    local.get 6
    i32.const 0
    i32.store offset=44
    local.get 6
    i32.const 24
    i32.add
    local.get 6
    i32.const 44
    i32.add
    local.get 6
    i32.load offset=8
    call 92
    local.get 6
    i32.const 24
    i32.add
    local.get 6
    i32.const 44
    i32.add
    local.get 6
    i64.load
    call 93
    local.get 6
    i32.const 24
    i32.add
    local.get 6
    i32.const 44
    i32.add
    local.get 6
    i32.load offset=12
    call 94
    local.get 7
    local.get 6
    i32.const 24
    i32.add
    i32.const 16
    call 15
    drop
    local.get 0
    local.get 7
    i64.const 0
    local.get 4
    local.get 5
    call 30
    drop
    local.get 6
    i32.const 48
    i32.add
    global.set 0)
  (func (;68;) (type 5) (result i32)
    (local i32)
    call 3
    local.tee 0
    call 10
    local.get 0)
  (func (;69;) (type 5) (result i32)
    call 63)
  (func (;70;) (type 5) (result i32)
    (local i32)
    call 3
    local.tee 0
    call 11
    local.get 0)
  (func (;71;) (type 14)
    block  ;; label = @1
      call 70
      call 68
      call 12
      i32.const 1
      i32.lt_s
      br_if 0 (;@1;)
      return
    end
    i32.const 1048652
    i32.const 36
    call 1
    unreachable)
  (func (;72;) (type 5) (result i32)
    call 73)
  (func (;73;) (type 5) (result i32)
    (local i32)
    i64.const 0
    call 2
    local.tee 0
    call 42
    local.get 0)
  (func (;74;) (type 14)
    block  ;; label = @1
      call 13
      i32.eqz
      br_if 0 (;@1;)
      i32.const 1048714
      i32.const 37
      call 1
      unreachable
    end
    call 73
    drop)
  (func (;75;) (type 8) (param i32) (result i32)
    (local i32)
    local.get 0
    call 3
    local.tee 1
    call 14
    drop
    local.get 1)
  (func (;76;) (type 0) (param i32 i32 i32 i32)
    (local i32)
    i32.const 1048688
    i32.const 23
    call 77
    local.tee 4
    local.get 0
    local.get 1
    call 15
    drop
    local.get 4
    i32.const 1048711
    i32.const 3
    call 15
    drop
    local.get 4
    local.get 2
    local.get 3
    call 15
    drop
    local.get 4
    call 16
    unreachable)
  (func (;77;) (type 1) (param i32 i32) (result i32)
    local.get 0
    local.get 1
    call 7)
  (func (;78;) (type 8) (param i32) (result i32)
    (local i32)
    local.get 0
    i64.const 0
    call 2
    local.tee 1
    call 17
    local.get 1)
  (func (;79;) (type 6) (param i32 i32 i32) (result i32)
    (local i64)
    block  ;; label = @1
      local.get 0
      call 18
      local.tee 3
      i64.const 4294967296
      i64.lt_u
      br_if 0 (;@1;)
      local.get 1
      local.get 2
      i32.const 1048576
      i32.const 14
      call 76
      unreachable
    end
    local.get 3
    i32.wrap_i64)
  (func (;80;) (type 8) (param i32) (result i32)
    block  ;; label = @1
      local.get 0
      call 75
      local.tee 0
      call 8
      i32.const 32
      i32.eq
      br_if 0 (;@1;)
      i32.const 1049194
      i32.const 4
      i32.const 1049069
      i32.const 16
      call 76
      unreachable
    end
    local.get 0)
  (func (;81;) (type 3) (param i32 i32)
    (local i32)
    local.get 1
    call 82
    local.set 2
    local.get 0
    i32.const 0
    i32.store offset=8
    local.get 0
    local.get 2
    i32.store offset=4
    local.get 0
    local.get 1
    i32.store)
  (func (;82;) (type 8) (param i32) (result i32)
    local.get 0
    call 57
    i32.const 2
    i32.shr_u)
  (func (;83;) (type 10) (param i32)
    block  ;; label = @1
      call 19
      local.get 0
      i32.ne
      br_if 0 (;@1;)
      return
    end
    i32.const 1048751
    i32.const 25
    call 1
    unreachable)
  (func (;84;) (type 0) (param i32 i32 i32 i32)
    local.get 0
    local.get 1
    local.get 2
    call 85
    local.get 0
    i32.const 1049676
    i32.const 1049680
    local.get 3
    select
    i32.const 4
    i32.const 5
    local.get 3
    select
    call 85)
  (func (;85;) (type 2) (param i32 i32 i32)
    local.get 0
    local.get 1
    local.get 2
    call 7
    call 60)
  (func (;86;) (type 2) (param i32 i32 i32)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 3
    global.set 0
    local.get 0
    call 87
    i32.const 1048776
    i32.const 14
    call 7
    call 88
    local.get 3
    call 3
    i32.store offset=12
    local.get 2
    local.get 3
    i32.const 12
    i32.add
    call 89
    local.get 0
    i32.const 24
    i32.add
    i32.load
    local.tee 2
    local.get 3
    i32.load offset=12
    call 60
    call 3
    drop
    local.get 2
    local.get 1
    call 51
    call 60
    i32.const -2
    local.set 0
    block  ;; label = @1
      loop  ;; label = @2
        local.get 0
        i32.eqz
        br_if 1 (;@1;)
        block  ;; label = @3
          local.get 0
          i32.const 1049493
          i32.add
          i32.load8_u
          local.tee 1
          i32.eqz
          br_if 0 (;@3;)
          local.get 2
          local.get 1
          i32.const 2
          i32.shl
          local.tee 1
          i32.const 1049912
          i32.add
          i32.load
          local.get 1
          i32.const 1049880
          i32.add
          i32.load
          call 85
        end
        local.get 0
        i32.const 1
        i32.add
        local.set 0
        br 0 (;@2;)
      end
    end
    local.get 3
    i32.const 16
    i32.add
    global.set 0)
  (func (;87;) (type 5) (result i32)
    i32.const 1048790
    call 96)
  (func (;88;) (type 2) (param i32 i32 i32)
    local.get 0
    local.get 1
    local.get 2
    call 3
    call 106)
  (func (;89;) (type 3) (param i32 i32)
    block  ;; label = @1
      local.get 0
      call 66
      br_if 0 (;@1;)
      local.get 1
      local.get 0
      call 51
      i32.store
      return
    end
    local.get 1
    i32.load
    i32.const 1049008
    i32.const 4
    call 126)
  (func (;90;) (type 16) (param i32 i32 i32 i32 i32)
    (local i32 i32 i32 i32 i32 i32 i32)
    global.get 0
    i32.const 80
    i32.sub
    local.tee 5
    global.set 0
    call 45
    drop
    local.get 4
    i32.load8_u offset=5
    local.set 6
    local.get 4
    i32.load8_u offset=4
    local.set 7
    local.get 4
    i32.load8_u offset=3
    local.set 8
    local.get 4
    i32.load8_u offset=2
    local.set 9
    local.get 4
    i32.load8_u offset=1
    local.set 10
    local.get 4
    i32.load8_u
    local.set 11
    local.get 5
    call 87
    i32.const 1048822
    i32.const 16
    call 7
    call 88
    local.get 5
    i32.const 32
    i32.add
    call 3
    i64.const 0
    local.get 1
    call 91
    local.get 5
    i32.const 56
    i32.add
    i32.const 8
    i32.add
    i64.const 0
    i64.store
    local.get 5
    i64.const 0
    i64.store offset=56
    local.get 5
    i32.const 0
    i32.store offset=76
    local.get 5
    i32.const 56
    i32.add
    local.get 5
    i32.const 76
    i32.add
    local.get 5
    i32.load offset=40
    call 92
    local.get 5
    i32.const 56
    i32.add
    local.get 5
    i32.const 76
    i32.add
    local.get 5
    i64.load offset=32
    call 93
    local.get 5
    i32.const 56
    i32.add
    local.get 5
    i32.const 76
    i32.add
    local.get 5
    i32.load offset=44
    call 94
    local.get 5
    i32.const 16
    i32.add
    local.tee 4
    i32.load
    local.get 5
    i32.const 56
    i32.add
    i32.const 16
    call 5
    drop
    local.get 0
    i32.const 24
    i32.add
    local.tee 1
    local.get 5
    i32.const 24
    i32.add
    i64.load
    i64.store
    local.get 0
    i32.const 16
    i32.add
    local.get 4
    i64.load
    i64.store
    local.get 0
    i32.const 8
    i32.add
    local.get 5
    i32.const 8
    i32.add
    i64.load
    i64.store
    local.get 0
    local.get 5
    i64.load
    i64.store
    local.get 1
    i32.load
    local.tee 4
    local.get 2
    call 95
    local.get 4
    local.get 3
    call 95
    local.get 4
    i32.const 1048838
    i32.const 9
    local.get 11
    call 84
    local.get 4
    i32.const 1048847
    i32.const 7
    local.get 10
    call 84
    local.get 4
    i32.const 1048854
    i32.const 8
    local.get 9
    call 84
    local.get 4
    i32.const 1048862
    i32.const 14
    local.get 8
    call 84
    local.get 4
    i32.const 1048876
    i32.const 10
    local.get 7
    call 84
    local.get 4
    i32.const 1048886
    i32.const 18
    local.get 6
    call 84
    local.get 5
    i32.const 80
    i32.add
    global.set 0)
  (func (;91;) (type 20) (param i32 i32 i64 i32)
    (local i32)
    i32.const 4
    local.set 4
    block  ;; label = @1
      local.get 3
      call 127
      i32.eqz
      br_if 0 (;@1;)
      local.get 1
      call 66
      br_if 0 (;@1;)
      block  ;; label = @2
        local.get 2
        i64.eqz
        i32.eqz
        br_if 0 (;@2;)
        i32.const 0
        local.set 4
        br 1 (;@1;)
      end
      i32.const 2
      i32.const 1
      local.get 3
      i64.const 1
      call 123
      i32.const 255
      i32.and
      select
      local.set 4
    end
    local.get 0
    local.get 1
    i32.store offset=8
    local.get 0
    local.get 4
    i32.store8 offset=16
    local.get 0
    local.get 3
    i32.store offset=12
    local.get 0
    local.get 2
    i64.store)
  (func (;92;) (type 2) (param i32 i32 i32)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 3
    global.set 0
    local.get 3
    local.get 2
    i32.const 24
    i32.shl
    local.get 2
    i32.const 8
    i32.shl
    i32.const 16711680
    i32.and
    i32.or
    local.get 2
    i32.const 8
    i32.shr_u
    i32.const 65280
    i32.and
    local.get 2
    i32.const 24
    i32.shr_u
    i32.or
    i32.or
    i32.store offset=12
    local.get 3
    local.get 1
    i32.load
    local.tee 2
    local.get 2
    i32.const 4
    i32.add
    local.tee 2
    local.get 0
    i32.const 16
    call 47
    local.get 3
    i32.load
    local.get 3
    i32.load offset=4
    local.get 3
    i32.const 12
    i32.add
    i32.const 4
    call 128
    local.get 1
    local.get 2
    i32.store
    local.get 3
    i32.const 16
    i32.add
    global.set 0)
  (func (;93;) (type 22) (param i32 i32 i64)
    (local i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 3
    global.set 0
    local.get 3
    local.get 2
    i64.const 56
    i64.shl
    local.get 2
    i64.const 40
    i64.shl
    i64.const 71776119061217280
    i64.and
    i64.or
    local.get 2
    i64.const 24
    i64.shl
    i64.const 280375465082880
    i64.and
    local.get 2
    i64.const 8
    i64.shl
    i64.const 1095216660480
    i64.and
    i64.or
    i64.or
    local.get 2
    i64.const 8
    i64.shr_u
    i64.const 4278190080
    i64.and
    local.get 2
    i64.const 24
    i64.shr_u
    i64.const 16711680
    i64.and
    i64.or
    local.get 2
    i64.const 40
    i64.shr_u
    i64.const 65280
    i64.and
    local.get 2
    i64.const 56
    i64.shr_u
    i64.or
    i64.or
    i64.or
    i64.store offset=8
    local.get 3
    local.get 1
    i32.load
    local.tee 4
    local.get 4
    i32.const 8
    i32.add
    local.tee 4
    local.get 0
    i32.const 16
    call 47
    local.get 3
    i32.load
    local.get 3
    i32.load offset=4
    local.get 3
    i32.const 8
    i32.add
    i32.const 8
    call 128
    local.get 1
    local.get 4
    i32.store
    local.get 3
    i32.const 16
    i32.add
    global.set 0)
  (func (;94;) (type 2) (param i32 i32 i32)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 3
    global.set 0
    local.get 3
    local.get 2
    i32.const 24
    i32.shl
    local.get 2
    i32.const 8
    i32.shl
    i32.const 16711680
    i32.and
    i32.or
    local.get 2
    i32.const 8
    i32.shr_u
    i32.const 65280
    i32.and
    local.get 2
    i32.const 24
    i32.shr_u
    i32.or
    i32.or
    i32.store offset=12
    local.get 3
    local.get 1
    i32.load
    local.tee 2
    local.get 2
    i32.const 4
    i32.add
    local.tee 2
    local.get 0
    i32.const 16
    call 47
    local.get 3
    i32.load
    local.get 3
    i32.load offset=4
    local.get 3
    i32.const 12
    i32.add
    i32.const 4
    call 128
    local.get 1
    local.get 2
    i32.store
    local.get 3
    i32.const 16
    i32.add
    global.set 0)
  (func (;95;) (type 3) (param i32 i32)
    call 3
    drop
    local.get 0
    local.get 1
    call 51
    call 60)
  (func (;96;) (type 8) (param i32) (result i32)
    local.get 0
    i32.const 32
    call 7)
  (func (;97;) (type 10) (param i32)
    (local i32 i32 i32 i32)
    global.get 0
    i32.const 48
    i32.sub
    local.tee 1
    global.set 0
    block  ;; label = @1
      local.get 0
      i32.load offset=16
      i32.const 1
      i32.ne
      br_if 0 (;@1;)
      call 98
      local.set 2
      local.get 1
      i32.const 16
      i32.add
      call 99
      local.get 1
      local.get 1
      i32.load8_u offset=20
      i32.store8 offset=28
      local.get 1
      local.get 1
      i32.load offset=16
      i32.store offset=24
      local.get 0
      i32.const 20
      i32.add
      i32.load
      local.get 1
      i32.const 24
      i32.add
      call 100
      local.get 0
      i32.const 24
      i32.add
      local.tee 3
      i32.load
      local.tee 4
      call 82
      local.get 1
      i32.const 24
      i32.add
      call 101
      local.get 1
      local.get 4
      call 57
      i32.store offset=40
      local.get 1
      i32.const 0
      i32.store offset=36
      local.get 1
      local.get 3
      i32.store offset=32
      block  ;; label = @2
        loop  ;; label = @3
          local.get 1
          i32.const 8
          i32.add
          local.get 1
          i32.const 32
          i32.add
          call 58
          local.get 1
          i32.load offset=8
          i32.eqz
          br_if 1 (;@2;)
          local.get 1
          i32.load offset=12
          local.get 1
          i32.const 24
          i32.add
          call 100
          br 0 (;@3;)
        end
      end
      local.get 2
      local.get 1
      i32.load offset=24
      local.get 1
      i32.load8_u offset=28
      call 102
    end
    local.get 0
    i32.load
    local.get 0
    i32.load offset=4
    local.get 0
    i32.load offset=8
    local.get 0
    i32.load offset=12
    call 103
    unreachable)
  (func (;98;) (type 5) (result i32)
    (local i32 i32)
    call 3
    local.tee 0
    call 20
    i32.const 1048951
    i32.const 10
    call 7
    local.tee 1
    local.get 0
    call 4
    drop
    local.get 1)
  (func (;99;) (type 10) (param i32)
    (local i32 i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 1
    global.set 0
    i32.const 0
    local.set 2
    block  ;; label = @1
      block  ;; label = @2
        i32.const 0
        i32.load8_u offset=1059948
        br_if 0 (;@2;)
        i32.const 1
        local.set 2
        i32.const 0
        i32.const 1
        i32.store8 offset=1059948
        i32.const 0
        i32.const 0
        i32.store offset=1049944
        local.get 1
        i32.const 8
        i32.add
        i32.const 0
        call 131
        local.get 1
        i32.load offset=8
        local.get 1
        i32.load offset=12
        i32.const 1049856
        i32.const 0
        call 128
        call 3
        local.set 3
        br 1 (;@1;)
      end
      i32.const 1049856
      i32.const 0
      call 7
      local.set 3
    end
    local.get 0
    local.get 2
    i32.store8 offset=4
    local.get 0
    local.get 3
    i32.store
    local.get 1
    i32.const 16
    i32.add
    global.set 0)
  (func (;100;) (type 3) (param i32 i32)
    (local i32 i32 i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    local.get 0
    call 57
    local.tee 3
    i32.const 24
    i32.shl
    local.get 3
    i32.const 8
    i32.shl
    i32.const 16711680
    i32.and
    i32.or
    local.get 3
    i32.const 8
    i32.shr_u
    i32.const 65280
    i32.and
    local.get 3
    i32.const 24
    i32.shr_u
    i32.or
    i32.or
    i32.store offset=12
    local.get 1
    local.get 2
    i32.const 12
    i32.add
    i32.const 4
    call 144
    local.get 1
    i32.load8_u offset=4
    local.set 3
    local.get 1
    i32.const 0
    i32.store8 offset=4
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            local.get 3
            i32.const 1
            i32.and
            local.tee 3
            i32.eqz
            br_if 0 (;@4;)
            local.get 0
            call 57
            local.set 4
            i32.const 10000
            i32.const 0
            i32.load offset=1049944
            local.tee 5
            i32.sub
            local.get 4
            i32.lt_u
            br_if 2 (;@2;)
            local.get 2
            local.get 5
            local.get 5
            local.get 4
            i32.add
            local.tee 4
            call 151
            local.get 0
            i32.const 0
            local.get 2
            i32.load
            local.get 2
            i32.load offset=4
            call 124
            drop
            i32.const 0
            local.get 4
            i32.store offset=1049944
            br 1 (;@3;)
          end
          local.get 1
          i32.load
          local.get 0
          call 125
        end
        local.get 1
        local.get 3
        i32.store8 offset=4
        br 1 (;@1;)
      end
      local.get 1
      call 133
      local.get 1
      i32.load
      local.get 0
      call 125
      local.get 1
      i32.load8_u offset=4
      local.set 0
      local.get 1
      local.get 3
      i32.store8 offset=4
      local.get 0
      i32.const 1
      i32.and
      i32.eqz
      br_if 0 (;@1;)
      i32.const 0
      i32.const 0
      i32.store offset=1049944
      i32.const 0
      i32.const 0
      i32.store8 offset=1059948
    end
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;101;) (type 3) (param i32 i32)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    local.get 0
    i32.const 24
    i32.shl
    local.get 0
    i32.const 8
    i32.shl
    i32.const 16711680
    i32.and
    i32.or
    local.get 0
    i32.const 8
    i32.shr_u
    i32.const 65280
    i32.and
    local.get 0
    i32.const 24
    i32.shr_u
    i32.or
    i32.or
    i32.store offset=12
    local.get 1
    local.get 2
    i32.const 12
    i32.add
    i32.const 4
    call 144
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;102;) (type 2) (param i32 i32 i32)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 3
    global.set 0
    local.get 3
    local.get 2
    i32.store8 offset=12
    local.get 3
    local.get 1
    i32.store offset=8
    local.get 3
    i32.const 8
    i32.add
    call 133
    local.get 3
    i32.load offset=8
    local.set 1
    block  ;; label = @1
      local.get 3
      i32.load8_u offset=12
      i32.eqz
      br_if 0 (;@1;)
      i32.const 0
      i32.const 0
      i32.store offset=1049944
      i32.const 0
      i32.const 0
      i32.store8 offset=1059948
    end
    local.get 0
    local.get 1
    call 26
    drop
    local.get 3
    i32.const 16
    i32.add
    global.set 0)
  (func (;103;) (type 0) (param i32 i32 i32 i32)
    local.get 0
    local.get 1
    local.get 2
    local.get 3
    call 29
    unreachable)
  (func (;104;) (type 3) (param i32 i32)
    (local i32 i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    i32.const 24
    i32.add
    local.get 1
    i32.const 24
    i32.add
    i64.load
    i64.store
    local.get 2
    i32.const 16
    i32.add
    local.get 1
    i32.const 16
    i32.add
    i64.load
    i64.store
    local.get 2
    i32.const 8
    i32.add
    local.get 1
    i32.const 8
    i32.add
    local.tee 3
    i64.load
    i64.store
    local.get 2
    local.get 1
    i64.load
    i64.store
    local.get 1
    local.get 2
    call 105
    local.get 0
    i32.const 0
    i32.store offset=16
    local.get 0
    local.get 1
    i64.load offset=20 align=4
    i64.store offset=8 align=4
    local.get 0
    local.get 3
    i64.load
    i64.store align=4
    local.get 2
    i32.const 32
    i32.add
    global.set 0)
  (func (;105;) (type 3) (param i32 i32)
    (local i32 i32 i32 i32 i64 i32 i32 i32 i32 i32 i64 i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 2
    global.set 0
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            local.get 1
            i32.load offset=16
            local.tee 3
            call 107
            br_table 1 (;@3;) 2 (;@2;) 0 (;@4;)
          end
          local.get 1
          i32.load offset=24
          local.set 4
          local.get 1
          i32.load offset=20
          local.set 5
          local.get 1
          i64.load
          local.set 6
          local.get 1
          i32.load offset=8
          local.set 7
          call 3
          local.set 8
          call 3
          local.set 1
          call 3
          drop
          local.get 1
          local.get 7
          call 51
          call 60
          local.get 3
          call 107
          local.set 7
          call 3
          local.tee 9
          local.get 7
          i64.extend_i32_u
          call 108
          local.get 1
          local.get 9
          call 60
          local.get 3
          call 8
          local.set 10
          local.get 2
          i32.const 16
          i32.add
          local.set 11
          i32.const 0
          local.set 7
          block  ;; label = @4
            loop  ;; label = @5
              local.get 7
              i32.const 16
              i32.add
              local.tee 9
              local.get 10
              i32.gt_u
              br_if 1 (;@4;)
              local.get 11
              i64.const 0
              i64.store
              local.get 2
              i64.const 0
              i64.store offset=8
              local.get 3
              local.get 7
              local.get 2
              i32.const 8
              i32.add
              i32.const 16
              call 62
              drop
              local.get 2
              i32.const 0
              i32.store offset=28
              local.get 2
              i32.const 8
              i32.add
              local.get 2
              i32.const 28
              i32.add
              call 109
              local.set 7
              local.get 2
              i32.const 8
              i32.add
              local.get 2
              i32.const 28
              i32.add
              call 110
              local.set 12
              local.get 2
              i32.const 8
              i32.add
              local.get 2
              i32.const 28
              i32.add
              call 111
              local.set 13
              local.get 2
              call 3
              i32.store offset=8
              local.get 7
              local.get 2
              i32.const 8
              i32.add
              call 89
              local.get 1
              local.get 2
              i32.load offset=8
              call 60
              call 3
              local.tee 7
              local.get 12
              call 108
              local.get 1
              local.get 7
              call 60
              call 3
              drop
              local.get 1
              local.get 13
              call 112
              call 60
              local.get 9
              local.set 7
              br 0 (;@5;)
            end
          end
          local.get 1
          local.get 5
          call 59
          call 63
          local.set 7
          call 45
          local.set 9
          i32.const 1048904
          i32.const 20
          call 7
          local.set 13
          local.get 1
          local.get 4
          call 113
          local.set 1
          local.get 0
          local.get 13
          i32.store offset=20
          local.get 0
          local.get 8
          i32.store offset=16
          local.get 0
          local.get 9
          i32.store offset=12
          local.get 0
          local.get 7
          i32.store offset=8
          local.get 0
          local.get 1
          i32.store offset=24
          local.get 0
          local.get 6
          i64.store
          br 2 (;@1;)
        end
        local.get 0
        local.get 1
        i64.load
        i64.store
        local.get 0
        i32.const 24
        i32.add
        local.get 1
        i32.const 24
        i32.add
        i64.load
        i64.store
        local.get 0
        i32.const 16
        i32.add
        local.get 1
        i32.const 16
        i32.add
        i64.load
        i64.store
        local.get 0
        i32.const 8
        i32.add
        local.get 1
        i32.const 8
        i32.add
        i64.load
        i64.store
        br 1 (;@1;)
      end
      local.get 1
      i32.load offset=28
      local.set 11
      local.get 1
      i32.load offset=24
      local.set 7
      local.get 1
      i32.load offset=20
      local.set 9
      local.get 1
      i32.load offset=12
      local.set 4
      local.get 1
      i32.load offset=8
      local.set 13
      local.get 1
      i64.load
      local.set 12
      local.get 2
      i32.const 16
      i32.add
      i64.const 0
      i64.store
      local.get 2
      i64.const 0
      i64.store offset=8
      local.get 3
      i32.const 0
      local.get 2
      i32.const 8
      i32.add
      i32.const 16
      call 62
      local.set 1
      local.get 2
      i32.const 0
      i32.store offset=28
      local.get 2
      i32.const 8
      i32.add
      local.get 2
      i32.const 28
      i32.add
      call 109
      local.set 10
      local.get 2
      i32.const 8
      i32.add
      local.get 2
      i32.const 28
      i32.add
      call 110
      local.set 6
      local.get 2
      i32.const 8
      i32.add
      local.get 2
      i32.const 28
      i32.add
      call 111
      local.set 5
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                local.get 1
                br_if 0 (;@6;)
                local.get 10
                call 66
                br_if 1 (;@5;)
                call 3
                local.set 3
                call 3
                local.tee 1
                local.get 10
                call 54
                local.get 6
                i64.eqz
                i32.eqz
                br_if 2 (;@4;)
                local.get 1
                local.get 5
                call 55
                local.get 1
                local.get 9
                call 56
                call 45
                local.set 9
                i32.const 1048939
                i32.const 12
                call 7
                local.set 10
                local.get 1
                local.get 7
                call 113
                local.set 1
                local.get 0
                local.get 10
                i32.store offset=20
                local.get 0
                local.get 3
                i32.store offset=16
                local.get 0
                local.get 9
                i32.store offset=12
                local.get 0
                local.get 13
                i32.store offset=8
                local.get 0
                local.get 1
                i32.store offset=24
                br 4 (;@2;)
              end
              local.get 0
              local.get 11
              i32.store offset=28
              local.get 0
              local.get 7
              i32.store offset=24
              local.get 0
              local.get 9
              i32.store offset=20
              local.get 0
              local.get 3
              i32.store offset=16
              local.get 0
              local.get 4
              i32.store offset=12
              br 2 (;@3;)
            end
            local.get 3
            i32.const 1049856
            i32.const 0
            call 5
            drop
            local.get 0
            local.get 11
            i32.store offset=28
            local.get 0
            local.get 7
            i32.store offset=24
            local.get 0
            local.get 9
            i32.store offset=20
            local.get 0
            local.get 3
            i32.store offset=16
            local.get 0
            local.get 5
            i32.store offset=12
            br 1 (;@3;)
          end
          local.get 1
          local.get 6
          call 114
          local.get 1
          local.get 5
          call 55
          call 3
          drop
          local.get 1
          local.get 13
          call 51
          call 60
          local.get 1
          local.get 9
          call 56
          call 63
          local.set 9
          call 45
          local.set 13
          i32.const 1048924
          i32.const 15
          call 7
          local.set 10
          local.get 1
          local.get 7
          call 113
          local.set 1
          local.get 0
          local.get 10
          i32.store offset=20
          local.get 0
          local.get 3
          i32.store offset=16
          local.get 0
          local.get 13
          i32.store offset=12
          local.get 0
          local.get 9
          i32.store offset=8
          local.get 0
          local.get 1
          i32.store offset=24
          br 1 (;@2;)
        end
        local.get 0
        local.get 13
        i32.store offset=8
      end
      local.get 0
      local.get 12
      i64.store
    end
    local.get 2
    i32.const 32
    i32.add
    global.set 0)
  (func (;106;) (type 0) (param i32 i32 i32 i32)
    (local i32 i32)
    call 3
    local.set 4
    call 45
    local.set 5
    local.get 0
    local.get 2
    i32.store offset=20
    local.get 0
    local.get 3
    i32.store offset=16
    local.get 0
    local.get 5
    i32.store offset=12
    local.get 0
    local.get 1
    i32.store offset=8
    local.get 0
    local.get 4
    i32.store offset=24
    local.get 0
    i64.const -1
    i64.store)
  (func (;107;) (type 8) (param i32) (result i32)
    local.get 0
    call 57
    i32.const 4
    i32.shr_u)
  (func (;108;) (type 23) (param i32 i64)
    (local i32 i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    local.get 1
    i64.const 56
    i64.shl
    local.get 1
    i64.const 40
    i64.shl
    i64.const 71776119061217280
    i64.and
    i64.or
    local.get 1
    i64.const 24
    i64.shl
    i64.const 280375465082880
    i64.and
    local.get 1
    i64.const 8
    i64.shl
    i64.const 1095216660480
    i64.and
    i64.or
    i64.or
    local.get 1
    i64.const 8
    i64.shr_u
    i64.const 4278190080
    i64.and
    local.get 1
    i64.const 24
    i64.shr_u
    i64.const 16711680
    i64.and
    i64.or
    local.get 1
    i64.const 40
    i64.shr_u
    i64.const 65280
    i64.and
    local.get 1
    i64.const 56
    i64.shr_u
    i64.or
    i64.or
    i64.or
    i64.store offset=8
    block  ;; label = @1
      block  ;; label = @2
        local.get 1
        i64.eqz
        i32.eqz
        br_if 0 (;@2;)
        i32.const 1049856
        local.set 3
        i32.const 0
        local.set 4
        br 1 (;@1;)
      end
      i32.const 0
      local.set 4
      loop  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              local.get 4
              i32.const 8
              i32.eq
              br_if 0 (;@5;)
              local.get 2
              i32.const 8
              i32.add
              local.get 4
              i32.add
              local.tee 3
              i32.load8_u
              i32.eqz
              br_if 2 (;@3;)
              local.get 4
              i32.const 9
              i32.ge_u
              br_if 1 (;@4;)
              i32.const 8
              local.get 4
              i32.sub
              local.set 4
              br 4 (;@1;)
            end
            i32.const 8
            i32.const 8
            call 148
            unreachable
          end
          local.get 4
          call 149
          unreachable
        end
        local.get 4
        i32.const 1
        i32.add
        local.set 4
        br 0 (;@2;)
      end
    end
    local.get 0
    local.get 3
    local.get 4
    call 126
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;109;) (type 1) (param i32 i32) (result i32)
    (local i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    i32.const 0
    i32.store offset=12
    local.get 2
    local.get 0
    local.get 1
    i32.load
    local.tee 3
    local.get 3
    i32.const 4
    i32.add
    local.tee 3
    call 129
    local.get 2
    i32.const 12
    i32.add
    i32.const 4
    local.get 2
    i32.load
    local.get 2
    i32.load offset=4
    call 128
    local.get 1
    local.get 3
    i32.store
    local.get 2
    i32.load offset=12
    local.set 1
    local.get 2
    i32.const 16
    i32.add
    global.set 0
    local.get 1
    i32.const 24
    i32.shl
    local.get 1
    i32.const 8
    i32.shl
    i32.const 16711680
    i32.and
    i32.or
    local.get 1
    i32.const 8
    i32.shr_u
    i32.const 65280
    i32.and
    local.get 1
    i32.const 24
    i32.shr_u
    i32.or
    i32.or)
  (func (;110;) (type 24) (param i32 i32) (result i64)
    (local i32 i32 i64)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    i64.const 0
    i64.store offset=8
    local.get 2
    local.get 0
    local.get 1
    i32.load
    local.tee 3
    local.get 3
    i32.const 8
    i32.add
    local.tee 3
    call 129
    local.get 2
    i32.const 8
    i32.add
    i32.const 8
    local.get 2
    i32.load
    local.get 2
    i32.load offset=4
    call 128
    local.get 1
    local.get 3
    i32.store
    local.get 2
    i64.load offset=8
    local.set 4
    local.get 2
    i32.const 16
    i32.add
    global.set 0
    local.get 4
    i64.const 56
    i64.shl
    local.get 4
    i64.const 40
    i64.shl
    i64.const 71776119061217280
    i64.and
    i64.or
    local.get 4
    i64.const 24
    i64.shl
    i64.const 280375465082880
    i64.and
    local.get 4
    i64.const 8
    i64.shl
    i64.const 1095216660480
    i64.and
    i64.or
    i64.or
    local.get 4
    i64.const 8
    i64.shr_u
    i64.const 4278190080
    i64.and
    local.get 4
    i64.const 24
    i64.shr_u
    i64.const 16711680
    i64.and
    i64.or
    local.get 4
    i64.const 40
    i64.shr_u
    i64.const 65280
    i64.and
    local.get 4
    i64.const 56
    i64.shr_u
    i64.or
    i64.or
    i64.or)
  (func (;111;) (type 1) (param i32 i32) (result i32)
    (local i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    i32.const 0
    i32.store offset=12
    local.get 2
    local.get 0
    local.get 1
    i32.load
    local.tee 3
    local.get 3
    i32.const 4
    i32.add
    local.tee 3
    call 129
    local.get 2
    i32.const 12
    i32.add
    i32.const 4
    local.get 2
    i32.load
    local.get 2
    i32.load offset=4
    call 128
    local.get 1
    local.get 3
    i32.store
    local.get 2
    i32.load offset=12
    local.set 1
    local.get 2
    i32.const 16
    i32.add
    global.set 0
    local.get 1
    i32.const 24
    i32.shl
    local.get 1
    i32.const 8
    i32.shl
    i32.const 16711680
    i32.and
    i32.or
    local.get 1
    i32.const 8
    i32.shr_u
    i32.const 65280
    i32.and
    local.get 1
    i32.const 24
    i32.shr_u
    i32.or
    i32.or)
  (func (;112;) (type 8) (param i32) (result i32)
    (local i32)
    call 3
    local.tee 1
    local.get 0
    call 23
    drop
    local.get 1)
  (func (;113;) (type 1) (param i32 i32) (result i32)
    local.get 0
    local.get 1
    call 4
    drop
    local.get 0)
  (func (;114;) (type 23) (param i32 i64)
    (local i32)
    call 3
    local.tee 2
    local.get 1
    call 108
    local.get 0
    local.get 2
    call 60)
  (func (;115;) (type 8) (param i32) (result i32)
    local.get 0
    call 116
    i64.const 0
    call 2
    local.tee 0
    call 21
    drop
    local.get 0)
  (func (;116;) (type 8) (param i32) (result i32)
    (local i32 i32 i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 1
    global.set 0
    local.get 0
    call 117
    local.set 2
    local.get 1
    i32.const 8
    i32.add
    local.get 0
    i32.load
    local.get 0
    i32.load offset=12
    local.tee 3
    local.get 2
    call 118
    block  ;; label = @1
      local.get 1
      i32.load offset=8
      i32.const 1
      i32.eq
      br_if 0 (;@1;)
      i32.const 1048961
      i32.const 15
      call 119
      unreachable
    end
    local.get 1
    i32.load offset=12
    local.set 4
    local.get 0
    local.get 3
    local.get 2
    i32.add
    i32.store offset=12
    local.get 1
    i32.const 16
    i32.add
    global.set 0
    local.get 4)
  (func (;117;) (type 8) (param i32) (result i32)
    (local i32 i64)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 1
    global.set 0
    local.get 1
    i32.const 0
    i32.store offset=12
    local.get 0
    local.get 1
    i32.const 12
    i32.add
    i32.const 4
    call 140
    local.get 1
    i32.const 12
    i32.add
    i32.const 4
    call 141
    local.set 2
    local.get 1
    i32.const 16
    i32.add
    global.set 0
    local.get 2
    i32.wrap_i64)
  (func (;118;) (type 0) (param i32 i32 i32 i32)
    (local i32)
    local.get 1
    local.get 2
    local.get 3
    call 3
    local.tee 4
    call 24
    local.set 1
    local.get 0
    local.get 4
    i32.store offset=4
    local.get 0
    local.get 1
    i32.eqz
    i32.store)
  (func (;119;) (type 3) (param i32 i32)
    (local i32)
    i32.const 1049047
    i32.const 22
    call 77
    local.tee 2
    local.get 0
    local.get 1
    call 15
    drop
    local.get 2
    call 16
    unreachable)
  (func (;120;) (type 3) (param i32 i32)
    (local i32)
    local.get 0
    local.get 1
    call 8
    local.tee 2
    i32.store offset=16
    local.get 0
    i32.const 0
    i32.store offset=12
    local.get 0
    i32.const 0
    i32.store8 offset=8
    local.get 0
    local.get 2
    i32.store offset=4
    local.get 0
    local.get 1
    i32.store)
  (func (;121;) (type 1) (param i32 i32) (result i32)
    local.get 0
    local.get 1
    call 122
    i32.const 255
    i32.and
    i32.eqz)
  (func (;122;) (type 1) (param i32 i32) (result i32)
    i32.const -1
    local.get 0
    local.get 1
    call 27
    local.tee 0
    i32.const 0
    i32.ne
    local.get 0
    i32.const 0
    i32.lt_s
    select)
  (func (;123;) (type 25) (param i32 i64) (result i32)
    block  ;; label = @1
      local.get 1
      i64.const 0
      i64.ne
      br_if 0 (;@1;)
      i32.const -1
      local.get 0
      call 22
      local.tee 0
      i32.const 0
      i32.ne
      local.get 0
      i32.const 0
      i32.lt_s
      select
      return
    end
    local.get 0
    local.get 1
    call 2
    call 122)
  (func (;124;) (type 12) (param i32 i32 i32 i32) (result i32)
    local.get 0
    local.get 1
    local.get 2
    local.get 3
    call 62)
  (func (;125;) (type 3) (param i32 i32)
    local.get 0
    local.get 1
    call 4
    drop)
  (func (;126;) (type 2) (param i32 i32 i32)
    local.get 0
    local.get 1
    local.get 2
    call 5
    drop)
  (func (;127;) (type 8) (param i32) (result i32)
    local.get 0
    i64.const 0
    call 123
    i32.const 255
    i32.and
    i32.const 0
    i32.ne)
  (func (;128;) (type 0) (param i32 i32 i32 i32)
    block  ;; label = @1
      local.get 1
      local.get 3
      i32.ne
      br_if 0 (;@1;)
      local.get 0
      local.get 2
      local.get 1
      call 200
      drop
      return
    end
    local.get 1
    local.get 3
    call 190
    unreachable)
  (func (;129;) (type 0) (param i32 i32 i32 i32)
    block  ;; label = @1
      block  ;; label = @2
        local.get 3
        local.get 2
        i32.lt_u
        br_if 0 (;@2;)
        local.get 3
        i32.const 16
        i32.gt_u
        br_if 1 (;@1;)
        local.get 0
        local.get 3
        local.get 2
        i32.sub
        i32.store offset=4
        local.get 0
        local.get 1
        local.get 2
        i32.add
        i32.store
        return
      end
      local.get 2
      local.get 3
      call 49
      unreachable
    end
    local.get 3
    i32.const 16
    call 48
    unreachable)
  (func (;130;) (type 12) (param i32 i32 i32 i32) (result i32)
    (local i32 i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 4
    global.set 0
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 0
          i32.load8_u offset=8
          br_if 0 (;@3;)
          local.get 0
          i32.load
          local.tee 5
          call 57
          local.set 6
          i32.const 0
          i32.load8_u offset=1059948
          br_if 1 (;@2;)
          local.get 6
          i32.const 10000
          i32.gt_u
          br_if 1 (;@2;)
          i32.const 0
          local.get 6
          i32.store offset=1049944
          i32.const 0
          i32.const 1
          i32.store8 offset=1059948
          local.get 4
          i32.const 8
          i32.add
          local.get 6
          call 131
          local.get 5
          i32.const 0
          local.get 4
          i32.load offset=8
          local.get 4
          i32.load offset=12
          call 124
          drop
          local.get 0
          i32.const 1
          i32.store8 offset=8
        end
        i32.const 1
        local.set 0
        local.get 3
        local.get 1
        i32.add
        local.tee 6
        i32.const 0
        i32.load offset=1049944
        i32.gt_u
        br_if 1 (;@1;)
        local.get 4
        local.get 1
        local.get 6
        call 132
        local.get 2
        local.get 3
        local.get 4
        i32.load
        local.get 4
        i32.load offset=4
        call 128
        i32.const 0
        local.set 0
        br 1 (;@1;)
      end
      local.get 0
      i32.const 0
      i32.store8 offset=8
      local.get 5
      local.get 1
      local.get 2
      local.get 3
      call 124
      local.set 0
    end
    local.get 4
    i32.const 16
    i32.add
    global.set 0
    local.get 0)
  (func (;131;) (type 3) (param i32 i32)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    i32.const 8
    i32.add
    i32.const 1049948
    i32.const 10000
    local.get 1
    call 157
    local.get 2
    i32.load offset=12
    local.set 1
    local.get 0
    local.get 2
    i32.load offset=8
    i32.store
    local.get 0
    local.get 1
    i32.store offset=4
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;132;) (type 2) (param i32 i32 i32)
    block  ;; label = @1
      block  ;; label = @2
        local.get 2
        local.get 1
        i32.lt_u
        br_if 0 (;@2;)
        local.get 2
        i32.const 10000
        i32.le_u
        br_if 1 (;@1;)
        local.get 2
        i32.const 10000
        call 48
        unreachable
      end
      local.get 1
      local.get 2
      call 49
      unreachable
    end
    local.get 0
    local.get 2
    local.get 1
    i32.sub
    i32.store offset=4
    local.get 0
    local.get 1
    i32.const 1049948
    i32.add
    i32.store)
  (func (;133;) (type 10) (param i32)
    (local i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 1
    global.set 0
    local.get 0
    i32.load8_u offset=4
    local.set 2
    local.get 0
    i32.const 0
    i32.store8 offset=4
    block  ;; label = @1
      local.get 2
      i32.const 1
      i32.and
      i32.eqz
      br_if 0 (;@1;)
      local.get 1
      i32.const 8
      i32.add
      i32.const 0
      i32.const 0
      i32.load offset=1049944
      call 132
      local.get 0
      i32.load
      local.get 1
      i32.load offset=8
      local.get 1
      i32.load offset=12
      call 15
      drop
      i32.const 0
      i32.const 0
      i32.store offset=1049944
      i32.const 0
      i32.const 0
      i32.store8 offset=1059948
    end
    local.get 1
    i32.const 16
    i32.add
    global.set 0)
  (func (;134;) (type 8) (param i32) (result i32)
    (local i32)
    local.get 0
    call 3
    local.tee 1
    call 25
    drop
    local.get 1)
  (func (;135;) (type 3) (param i32 i32)
    local.get 0
    local.get 1
    call 26
    drop)
  (func (;136;) (type 8) (param i32) (result i32)
    local.get 0
    call 137
    call 138)
  (func (;137;) (type 8) (param i32) (result i32)
    local.get 0
    call 134)
  (func (;138;) (type 8) (param i32) (result i32)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 1
    global.set 0
    block  ;; label = @1
      local.get 0
      call 8
      i32.const 4
      i32.ne
      br_if 0 (;@1;)
      local.get 1
      i32.const 0
      i32.store offset=12
      local.get 0
      i32.const 0
      local.get 1
      i32.const 12
      i32.add
      i32.const 4
      call 62
      drop
      local.get 1
      i32.const 12
      i32.add
      i32.const 4
      i32.const 1049008
      i32.const 4
      call 153
      i32.eqz
      br_if 0 (;@1;)
      local.get 0
      i32.const 1049856
      i32.const 0
      call 5
      drop
    end
    local.get 1
    i32.const 16
    i32.add
    global.set 0
    local.get 0)
  (func (;139;) (type 3) (param i32 i32)
    (local i32 i32 i32 i32 i32 i64)
    global.get 0
    i32.const 80048
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    i32.const 8
    i32.add
    local.get 1
    call 134
    call 120
    local.get 2
    i32.const 8
    i32.add
    call 115
    local.set 3
    local.get 2
    i32.const 8
    i32.add
    call 115
    local.set 4
    block  ;; label = @1
      block  ;; label = @2
        local.get 2
        i32.const 8
        i32.add
        call 117
        local.tee 5
        i32.const 10000
        i32.gt_u
        br_if 0 (;@2;)
        i32.const 0
        local.set 1
        local.get 2
        i32.const 40040
        i32.add
        local.set 6
        block  ;; label = @3
          loop  ;; label = @4
            local.get 2
            local.get 1
            i32.store offset=80040
            local.get 5
            local.get 1
            i32.eq
            br_if 1 (;@3;)
            local.get 6
            local.get 2
            i32.const 8
            i32.add
            call 117
            i32.store
            local.get 6
            i32.const 4
            i32.add
            local.set 6
            local.get 1
            i32.const 1
            i32.add
            local.set 1
            br 0 (;@4;)
          end
        end
        local.get 2
        i32.const 32
        i32.add
        local.get 2
        i32.const 40040
        i32.add
        i32.const 40004
        call 200
        drop
        local.get 2
        i32.const 8
        i32.add
        call 116
        local.set 1
        local.get 2
        i32.const 8
        i32.add
        call 116
        local.set 6
        local.get 2
        i32.const 8
        i32.add
        call 116
        local.set 5
        local.get 2
        i64.const 0
        i64.store offset=40040
        local.get 2
        i32.const 8
        i32.add
        local.get 2
        i32.const 40040
        i32.add
        i32.const 8
        call 140
        local.get 2
        i32.const 40040
        i32.add
        i32.const 8
        call 141
        local.set 7
        local.get 2
        i32.load offset=24
        local.get 2
        i32.load offset=20
        i32.ne
        br_if 1 (;@1;)
        local.get 0
        i32.const 16
        i32.add
        local.get 2
        i32.const 32
        i32.add
        i32.const 40004
        call 200
        drop
        block  ;; label = @3
          local.get 2
          i32.load8_u offset=16
          i32.eqz
          br_if 0 (;@3;)
          i32.const 0
          i32.const 0
          i32.store offset=1049944
          i32.const 0
          i32.const 0
          i32.store8 offset=1059948
        end
        local.get 0
        local.get 5
        i32.store offset=40028
        local.get 0
        local.get 6
        i32.store offset=40024
        local.get 0
        local.get 1
        i32.store offset=40020
        local.get 0
        local.get 4
        i32.store offset=12
        local.get 0
        local.get 3
        i32.store offset=8
        local.get 0
        local.get 7
        i64.store
        local.get 2
        i32.const 80048
        i32.add
        global.set 0
        return
      end
      i32.const 1049012
      i32.const 17
      call 119
      unreachable
    end
    i32.const 1048576
    i32.const 14
    call 119
    unreachable)
  (func (;140;) (type 2) (param i32 i32 i32)
    block  ;; label = @1
      local.get 0
      local.get 0
      i32.load offset=12
      local.get 1
      local.get 2
      call 130
      i32.eqz
      br_if 0 (;@1;)
      i32.const 1048961
      i32.const 15
      call 119
      unreachable
    end
    local.get 0
    local.get 0
    i32.load offset=12
    local.get 2
    i32.add
    i32.store offset=12)
  (func (;141;) (type 24) (param i32 i32) (result i64)
    (local i64)
    i64.const 0
    local.set 2
    block  ;; label = @1
      loop  ;; label = @2
        local.get 1
        i32.eqz
        br_if 1 (;@1;)
        local.get 1
        i32.const -1
        i32.add
        local.set 1
        local.get 2
        i64.const 8
        i64.shl
        local.get 0
        i64.load8_u
        i64.or
        local.set 2
        local.get 0
        i32.const 1
        i32.add
        local.set 0
        br 0 (;@2;)
      end
    end
    local.get 2)
  (func (;142;) (type 3) (param i32 i32)
    (local i32 i32 i32 i64)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    i32.const 8
    i32.add
    call 99
    local.get 2
    local.get 2
    i32.load8_u offset=12
    i32.store8 offset=20
    local.get 2
    local.get 2
    i32.load offset=8
    i32.store offset=16
    local.get 1
    i32.load offset=8
    local.get 2
    i32.const 16
    i32.add
    call 143
    local.get 1
    i32.load offset=12
    local.get 2
    i32.const 16
    i32.add
    call 143
    local.get 1
    i32.const 16
    i32.add
    local.set 3
    local.get 1
    i32.const 40016
    i32.add
    i32.load
    local.tee 4
    local.get 2
    i32.const 16
    i32.add
    call 101
    local.get 4
    i32.const 2
    i32.shl
    local.set 4
    block  ;; label = @1
      loop  ;; label = @2
        local.get 4
        i32.eqz
        br_if 1 (;@1;)
        local.get 3
        i32.load
        local.get 2
        i32.const 16
        i32.add
        call 101
        local.get 4
        i32.const -4
        i32.add
        local.set 4
        local.get 3
        i32.const 4
        i32.add
        local.set 3
        br 0 (;@2;)
      end
    end
    local.get 1
    i32.load offset=40020
    local.get 2
    i32.const 16
    i32.add
    call 100
    local.get 1
    i32.load offset=40024
    local.get 2
    i32.const 16
    i32.add
    call 100
    local.get 1
    i32.load offset=40028
    local.get 2
    i32.const 16
    i32.add
    call 100
    local.get 2
    local.get 1
    i64.load
    local.tee 5
    i64.const 56
    i64.shl
    local.get 5
    i64.const 40
    i64.shl
    i64.const 71776119061217280
    i64.and
    i64.or
    local.get 5
    i64.const 24
    i64.shl
    i64.const 280375465082880
    i64.and
    local.get 5
    i64.const 8
    i64.shl
    i64.const 1095216660480
    i64.and
    i64.or
    i64.or
    local.get 5
    i64.const 8
    i64.shr_u
    i64.const 4278190080
    i64.and
    local.get 5
    i64.const 24
    i64.shr_u
    i64.const 16711680
    i64.and
    i64.or
    local.get 5
    i64.const 40
    i64.shr_u
    i64.const 65280
    i64.and
    local.get 5
    i64.const 56
    i64.shr_u
    i64.or
    i64.or
    i64.or
    i64.store offset=24
    local.get 2
    i32.const 16
    i32.add
    local.get 2
    i32.const 24
    i32.add
    i32.const 8
    call 144
    local.get 0
    local.get 2
    i32.load offset=16
    local.get 2
    i32.load8_u offset=20
    call 102
    local.get 2
    i32.const 32
    i32.add
    global.set 0)
  (func (;143;) (type 3) (param i32 i32)
    local.get 0
    call 112
    local.get 1
    call 100)
  (func (;144;) (type 2) (param i32 i32 i32)
    (local i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 3
    global.set 0
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 0
          i32.load8_u offset=4
          i32.eqz
          br_if 0 (;@3;)
          i32.const 10000
          i32.const 0
          i32.load offset=1049944
          local.tee 4
          i32.sub
          local.get 2
          i32.lt_u
          br_if 1 (;@2;)
          local.get 3
          i32.const 8
          i32.add
          local.get 4
          local.get 4
          local.get 2
          i32.add
          local.tee 0
          call 151
          local.get 3
          i32.load offset=8
          local.get 3
          i32.load offset=12
          local.get 1
          local.get 2
          call 128
          i32.const 0
          local.get 0
          i32.store offset=1049944
          br 2 (;@1;)
        end
        local.get 0
        i32.load
        local.get 1
        local.get 2
        call 15
        drop
        br 1 (;@1;)
      end
      local.get 0
      call 133
      local.get 0
      i32.load
      local.get 1
      local.get 2
      call 15
      drop
    end
    local.get 3
    i32.const 16
    i32.add
    global.set 0)
  (func (;145;) (type 3) (param i32 i32)
    block  ;; label = @1
      local.get 1
      call 66
      br_if 0 (;@1;)
      local.get 0
      local.get 1
      call 135
      return
    end
    local.get 0
    i32.const 1049008
    i32.const 4
    call 7
    call 26
    drop)
  (func (;146;) (type 8) (param i32) (result i32)
    (local i32 i64)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 1
    global.set 0
    local.get 1
    i32.const 0
    i32.store offset=4
    local.get 1
    i32.const 25
    i32.store offset=12
    local.get 1
    i32.const 1048590
    i32.store offset=8
    block  ;; label = @1
      local.get 0
      local.get 0
      i32.load offset=12
      local.get 1
      i32.const 4
      i32.add
      i32.const 4
      call 130
      i32.eqz
      br_if 0 (;@1;)
      local.get 1
      i32.const 8
      i32.add
      i32.const 1048961
      i32.const 15
      call 147
      unreachable
    end
    local.get 0
    local.get 0
    i32.load offset=12
    i32.const 4
    i32.add
    i32.store offset=12
    local.get 1
    i32.const 4
    i32.add
    i32.const 4
    call 141
    local.set 2
    local.get 1
    i32.const 16
    i32.add
    global.set 0
    local.get 2
    i32.wrap_i64)
  (func (;147;) (type 2) (param i32 i32 i32)
    local.get 0
    i32.load
    local.get 0
    i32.load offset=4
    call 77
    local.tee 0
    local.get 1
    local.get 2
    call 15
    drop
    local.get 0
    call 16
    unreachable)
  (func (;148;) (type 3) (param i32 i32)
    call 164
    unreachable)
  (func (;149;) (type 10) (param i32)
    local.get 0
    call 181
    unreachable)
  (func (;150;) (type 8) (param i32) (result i32)
    (local i32 i32 i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 1
    global.set 0
    local.get 0
    call 146
    local.set 2
    local.get 1
    i32.const 25
    i32.store offset=12
    local.get 1
    i32.const 1048590
    i32.store offset=8
    local.get 1
    local.get 0
    i32.load
    local.get 0
    i32.load offset=12
    local.tee 3
    local.get 2
    call 118
    block  ;; label = @1
      local.get 1
      i32.load
      i32.const 1
      i32.eq
      br_if 0 (;@1;)
      local.get 1
      i32.const 8
      i32.add
      i32.const 1048961
      i32.const 15
      call 147
      unreachable
    end
    local.get 1
    i32.load offset=4
    local.set 4
    local.get 0
    local.get 3
    local.get 2
    i32.add
    i32.store offset=12
    local.get 1
    i32.const 16
    i32.add
    global.set 0
    local.get 4)
  (func (;151;) (type 2) (param i32 i32 i32)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 3
    global.set 0
    local.get 3
    i32.const 8
    i32.add
    local.get 1
    local.get 2
    i32.const 1049948
    i32.const 10000
    call 47
    local.get 3
    i32.load offset=12
    local.set 1
    local.get 0
    local.get 3
    i32.load offset=8
    i32.store
    local.get 0
    local.get 1
    i32.store offset=4
    local.get 3
    i32.const 16
    i32.add
    global.set 0)
  (func (;152;) (type 8) (param i32) (result i32)
    (local i32)
    block  ;; label = @1
      local.get 0
      i32.load
      local.tee 1
      local.get 0
      i32.load offset=4
      i32.ge_s
      br_if 0 (;@1;)
      local.get 0
      local.get 1
      i32.const 1
      i32.add
      i32.store
      local.get 1
      return
    end
    i32.const 1049166
    i32.const 6
    i32.const 1049085
    i32.const 17
    call 76
    unreachable)
  (func (;153;) (type 12) (param i32 i32 i32 i32) (result i32)
    (local i32)
    i32.const 0
    local.set 4
    block  ;; label = @1
      local.get 1
      local.get 3
      i32.ne
      br_if 0 (;@1;)
      local.get 0
      local.get 2
      local.get 1
      call 199
      i32.eqz
      local.set 4
    end
    local.get 4)
  (func (;154;) (type 5) (result i32)
    i32.const 1049102
    i32.const 10
    call 7)
  (func (;155;) (type 5) (result i32)
    i32.const 1049112
    i32.const 15
    call 7)
  (func (;156;) (type 3) (param i32 i32)
    (local i32 i32 i32)
    i32.const 0
    local.set 2
    block  ;; label = @1
      block  ;; label = @2
        local.get 1
        i32.load8_u offset=8
        i32.eqz
        br_if 0 (;@2;)
        br 1 (;@1;)
      end
      local.get 1
      i32.load
      local.tee 3
      local.get 1
      i32.load offset=4
      local.tee 4
      i32.gt_u
      br_if 0 (;@1;)
      block  ;; label = @2
        local.get 3
        local.get 4
        i32.lt_u
        br_if 0 (;@2;)
        i32.const 1
        local.set 2
        local.get 1
        i32.const 1
        i32.store8 offset=8
        br 1 (;@1;)
      end
      i32.const 1
      local.set 2
      local.get 1
      local.get 3
      i32.const 1
      i32.add
      i32.store
    end
    local.get 0
    local.get 3
    i32.store offset=4
    local.get 0
    local.get 2
    i32.store)
  (func (;157;) (type 0) (param i32 i32 i32 i32)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 4
    global.set 0
    local.get 4
    i32.const 8
    i32.add
    i32.const 0
    local.get 3
    local.get 1
    local.get 2
    call 47
    local.get 4
    i32.load offset=12
    local.set 3
    local.get 0
    local.get 4
    i32.load offset=8
    i32.store
    local.get 0
    local.get 3
    i32.store offset=4
    local.get 4
    i32.const 16
    i32.add
    global.set 0)
  (func (;158;) (type 10) (param i32)
    (local i32)
    i32.const 1049144
    i32.const 14
    call 7
    local.set 1
    local.get 0
    call 3
    i32.store offset=4
    local.get 0
    local.get 1
    i32.store)
  (func (;159;) (type 14)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i64)
    global.get 0
    i32.const 80080
    i32.sub
    local.tee 0
    global.set 0
    call 154
    call 136
    local.set 1
    local.get 0
    call 155
    call 139
    local.get 0
    i32.const 40040
    i32.add
    call 155
    call 139
    call 3
    local.set 2
    local.get 0
    i32.const 80056
    i32.add
    i32.load
    local.set 3
    local.get 2
    i32.const 4
    call 32
    drop
    local.get 0
    i32.const 0
    i32.store offset=80076
    local.get 2
    i32.const 0
    local.get 0
    i32.const 80076
    i32.add
    i32.const 4
    call 62
    drop
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  local.get 3
                  i32.eqz
                  br_if 0 (;@7;)
                  local.get 0
                  i32.const 40056
                  i32.add
                  local.tee 4
                  local.get 0
                  i32.load offset=80076
                  local.tee 2
                  i32.const 24
                  i32.shl
                  local.get 2
                  i32.const 8
                  i32.shl
                  i32.const 16711680
                  i32.and
                  i32.or
                  local.get 2
                  i32.const 8
                  i32.shr_u
                  i32.const 65280
                  i32.and
                  local.get 2
                  i32.const 24
                  i32.shr_u
                  i32.or
                  i32.or
                  local.get 3
                  i32.rem_u
                  i32.const 2
                  i32.shl
                  i32.add
                  local.tee 5
                  i32.load
                  local.set 2
                  local.get 5
                  local.get 3
                  i32.const 2
                  i32.shl
                  local.get 4
                  i32.add
                  i32.const -4
                  i32.add
                  local.tee 3
                  i32.load
                  i32.store
                  local.get 3
                  local.get 2
                  i32.store
                  local.get 0
                  i32.load offset=80056
                  local.tee 3
                  i32.eqz
                  br_if 1 (;@6;)
                  local.get 0
                  local.get 3
                  i32.const -1
                  i32.add
                  i32.store offset=80056
                  call 155
                  local.get 0
                  i32.const 40040
                  i32.add
                  call 142
                  local.get 4
                  i32.const 0
                  i32.store
                  local.get 0
                  i32.const 40040
                  i32.add
                  i32.const 8
                  i32.add
                  i64.const 0
                  i64.store
                  local.get 0
                  i64.const 0
                  i64.store offset=40040
                  block  ;; label = @8
                    block  ;; label = @9
                      block  ;; label = @10
                        local.get 2
                        i32.eqz
                        br_if 0 (;@10;)
                        i32.const 19
                        local.set 3
                        loop  ;; label = @11
                          local.get 2
                          i32.eqz
                          br_if 2 (;@9;)
                          block  ;; label = @12
                            local.get 3
                            i32.const 19
                            i32.gt_u
                            br_if 0 (;@12;)
                            local.get 0
                            i32.const 40040
                            i32.add
                            local.get 3
                            i32.add
                            local.get 2
                            i32.const 10
                            i32.rem_u
                            i32.const 48
                            i32.or
                            i32.store8
                            local.get 3
                            i32.const -1
                            i32.add
                            local.set 3
                            local.get 2
                            i32.const 10
                            i32.div_u
                            local.set 2
                            br 1 (;@11;)
                          end
                        end
                        i32.const -1
                        i32.const 20
                        call 148
                        unreachable
                      end
                      local.get 0
                      i32.const 48
                      i32.store8 offset=40059
                      i32.const 19
                      local.set 5
                      br 1 (;@8;)
                    end
                    local.get 3
                    i32.const 1
                    i32.add
                    local.set 5
                  end
                  i32.const 1
                  local.set 3
                  local.get 0
                  i32.load offset=40020
                  local.set 2
                  call 3
                  local.set 4
                  local.get 2
                  call 51
                  local.tee 2
                  i32.const 1049252
                  i32.const 1
                  call 15
                  drop
                  local.get 2
                  call 51
                  local.set 6
                  local.get 2
                  local.get 0
                  i32.const 40040
                  i32.add
                  local.get 5
                  i32.add
                  local.tee 7
                  i32.const 20
                  local.get 5
                  i32.sub
                  local.tee 5
                  call 15
                  drop
                  local.get 2
                  call 51
                  local.set 8
                  local.get 2
                  i32.const 1049253
                  i32.const 4
                  call 15
                  drop
                  local.get 4
                  local.get 2
                  call 60
                  local.get 8
                  i32.const 1049257
                  i32.const 5
                  call 15
                  drop
                  local.get 4
                  local.get 8
                  call 60
                  local.get 6
                  i32.const 1049262
                  i32.const 15
                  call 15
                  drop
                  local.get 4
                  local.get 6
                  call 60
                  local.get 0
                  local.get 4
                  i32.store offset=40036
                  local.get 0
                  i32.load offset=40028
                  call 51
                  local.tee 6
                  i32.const 1049292
                  i32.const 3
                  call 15
                  drop
                  local.get 6
                  local.get 7
                  local.get 5
                  call 15
                  drop
                  local.get 0
                  i32.load offset=40024
                  local.set 8
                  call 3
                  local.tee 2
                  i32.const 1049277
                  i32.const 5
                  call 15
                  drop
                  local.get 2
                  i32.const 1049295
                  i32.const 20
                  call 15
                  drop
                  local.get 2
                  i32.const 1049282
                  i32.const 10
                  call 15
                  drop
                  block  ;; label = @8
                    local.get 8
                    call 8
                    local.tee 4
                    i32.eqz
                    br_if 0 (;@8;)
                    block  ;; label = @9
                      block  ;; label = @10
                        local.get 4
                        i32.const 3
                        i32.add
                        i32.const 2
                        i32.shr_u
                        local.tee 3
                        i32.const -1
                        i32.add
                        local.tee 9
                        i32.const 256
                        i32.lt_u
                        br_if 0 (;@10;)
                        local.get 0
                        i32.const 0
                        i32.load offset=1059984
                        i32.store offset=80076
                        local.get 3
                        i32.const 1
                        local.get 0
                        i32.const 80076
                        i32.add
                        i32.const 1049856
                        i32.const 1049856
                        call 160
                        local.set 3
                        i32.const 0
                        local.get 0
                        i32.load offset=80076
                        i32.store offset=1059984
                        br 1 (;@9;)
                      end
                      local.get 0
                      i32.const 1059984
                      i32.store offset=80072
                      local.get 0
                      local.get 9
                      i32.const 2
                      i32.shl
                      i32.const 1059988
                      i32.add
                      local.tee 9
                      i32.load
                      i32.store offset=80076
                      local.get 3
                      i32.const 1
                      local.get 0
                      i32.const 80076
                      i32.add
                      local.get 0
                      i32.const 80072
                      i32.add
                      i32.const 1049832
                      call 160
                      local.set 3
                      local.get 9
                      local.get 0
                      i32.load offset=80076
                      i32.store
                    end
                    local.get 4
                    i32.const 1
                    i32.lt_s
                    br_if 0 (;@8;)
                    local.get 8
                    local.get 3
                    call 33
                    drop
                  end
                  local.get 2
                  local.get 3
                  local.get 4
                  call 15
                  drop
                  local.get 4
                  i32.eqz
                  br_if 6 (;@1;)
                  block  ;; label = @8
                    local.get 4
                    i32.const 3
                    i32.add
                    i32.const 2
                    i32.shr_u
                    i32.const -1
                    i32.add
                    local.tee 8
                    i32.const 256
                    i32.lt_u
                    br_if 0 (;@8;)
                    local.get 3
                    i32.const 0
                    i32.store
                    local.get 3
                    i32.const -8
                    i32.add
                    local.tee 4
                    local.get 4
                    i32.load
                    local.tee 9
                    i32.const -2
                    i32.and
                    i32.store
                    i32.const 0
                    i32.load offset=1059984
                    local.set 8
                    i32.const 1049856
                    call 161
                    i32.eqz
                    br_if 5 (;@3;)
                    block  ;; label = @9
                      block  ;; label = @10
                        local.get 3
                        i32.const -4
                        i32.add
                        local.tee 10
                        i32.load
                        i32.const -4
                        i32.and
                        local.tee 11
                        i32.eqz
                        br_if 0 (;@10;)
                        local.get 11
                        i32.load8_u
                        i32.const 1
                        i32.and
                        br_if 0 (;@10;)
                        local.get 11
                        local.set 3
                        block  ;; label = @11
                          block  ;; label = @12
                            local.get 9
                            i32.const -4
                            i32.and
                            local.tee 12
                            i32.eqz
                            br_if 0 (;@12;)
                            local.get 11
                            local.set 3
                            i32.const 0
                            local.get 12
                            local.get 9
                            i32.const 2
                            i32.and
                            select
                            local.tee 9
                            i32.eqz
                            br_if 0 (;@12;)
                            local.get 9
                            local.get 9
                            i32.load offset=4
                            i32.const 3
                            i32.and
                            local.get 11
                            i32.or
                            i32.store offset=4
                            local.get 10
                            i32.load
                            local.tee 9
                            i32.const -4
                            i32.and
                            local.tee 3
                            i32.eqz
                            br_if 1 (;@11;)
                          end
                          local.get 3
                          local.get 3
                          i32.load
                          i32.const 3
                          i32.and
                          local.get 4
                          i32.load
                          i32.const -4
                          i32.and
                          i32.or
                          i32.store
                          local.get 10
                          i32.load
                          local.set 9
                        end
                        local.get 10
                        local.get 9
                        i32.const 3
                        i32.and
                        i32.store
                        local.get 4
                        local.get 4
                        i32.load
                        local.tee 3
                        i32.const 3
                        i32.and
                        i32.store
                        local.get 3
                        i32.const 2
                        i32.and
                        i32.eqz
                        br_if 1 (;@9;)
                        local.get 11
                        local.get 11
                        i32.load
                        i32.const 2
                        i32.or
                        i32.store
                        br 1 (;@9;)
                      end
                      local.get 9
                      i32.const -4
                      i32.and
                      local.tee 11
                      i32.eqz
                      br_if 6 (;@3;)
                      i32.const 0
                      local.get 11
                      local.get 9
                      i32.const 2
                      i32.and
                      select
                      local.tee 9
                      i32.eqz
                      br_if 6 (;@3;)
                      local.get 9
                      i32.load8_u
                      i32.const 1
                      i32.and
                      br_if 6 (;@3;)
                      local.get 3
                      local.get 9
                      i32.load offset=8
                      i32.const -4
                      i32.and
                      i32.store
                      local.get 9
                      local.get 4
                      i32.const 1
                      i32.or
                      i32.store offset=8
                    end
                    local.get 8
                    local.set 4
                    br 6 (;@2;)
                  end
                  local.get 0
                  i32.const 1059984
                  i32.store offset=80076
                  local.get 3
                  i32.const 0
                  i32.store
                  local.get 3
                  i32.const -8
                  i32.add
                  local.tee 4
                  local.get 4
                  i32.load
                  local.tee 11
                  i32.const -2
                  i32.and
                  i32.store
                  local.get 8
                  i32.const 2
                  i32.shl
                  i32.const 1059988
                  i32.add
                  local.tee 8
                  i32.load
                  local.set 9
                  local.get 0
                  i32.const 80076
                  i32.add
                  call 162
                  i32.eqz
                  br_if 2 (;@5;)
                  block  ;; label = @8
                    block  ;; label = @9
                      local.get 3
                      i32.const -4
                      i32.add
                      local.tee 12
                      i32.load
                      i32.const -4
                      i32.and
                      local.tee 10
                      i32.eqz
                      br_if 0 (;@9;)
                      local.get 10
                      i32.load8_u
                      i32.const 1
                      i32.and
                      br_if 0 (;@9;)
                      local.get 10
                      local.set 3
                      block  ;; label = @10
                        block  ;; label = @11
                          local.get 11
                          i32.const -4
                          i32.and
                          local.tee 13
                          i32.eqz
                          br_if 0 (;@11;)
                          local.get 10
                          local.set 3
                          i32.const 0
                          local.get 13
                          local.get 11
                          i32.const 2
                          i32.and
                          select
                          local.tee 11
                          i32.eqz
                          br_if 0 (;@11;)
                          local.get 11
                          local.get 11
                          i32.load offset=4
                          i32.const 3
                          i32.and
                          local.get 10
                          i32.or
                          i32.store offset=4
                          local.get 12
                          i32.load
                          local.tee 11
                          i32.const -4
                          i32.and
                          local.tee 3
                          i32.eqz
                          br_if 1 (;@10;)
                        end
                        local.get 3
                        local.get 3
                        i32.load
                        i32.const 3
                        i32.and
                        local.get 4
                        i32.load
                        i32.const -4
                        i32.and
                        i32.or
                        i32.store
                        local.get 12
                        i32.load
                        local.set 11
                      end
                      local.get 12
                      local.get 11
                      i32.const 3
                      i32.and
                      i32.store
                      local.get 4
                      local.get 4
                      i32.load
                      local.tee 3
                      i32.const 3
                      i32.and
                      i32.store
                      local.get 3
                      i32.const 2
                      i32.and
                      i32.eqz
                      br_if 1 (;@8;)
                      local.get 10
                      local.get 10
                      i32.load
                      i32.const 2
                      i32.or
                      i32.store
                      br 1 (;@8;)
                    end
                    local.get 11
                    i32.const -4
                    i32.and
                    local.tee 10
                    i32.eqz
                    br_if 3 (;@5;)
                    i32.const 0
                    local.get 10
                    local.get 11
                    i32.const 2
                    i32.and
                    select
                    local.tee 11
                    i32.eqz
                    br_if 3 (;@5;)
                    local.get 11
                    i32.load8_u
                    i32.const 1
                    i32.and
                    br_if 3 (;@5;)
                    local.get 3
                    local.get 11
                    i32.load offset=8
                    i32.const -4
                    i32.and
                    i32.store
                    local.get 11
                    local.get 4
                    i32.const 1
                    i32.or
                    i32.store offset=8
                  end
                  local.get 9
                  local.set 4
                  br 3 (;@4;)
                end
                call 163
                unreachable
              end
              call 164
              unreachable
            end
            local.get 3
            local.get 9
            i32.store
          end
          local.get 8
          local.get 4
          i32.store
          br 2 (;@1;)
        end
        local.get 3
        local.get 8
        i32.store
      end
      i32.const 0
      local.get 4
      i32.store offset=1059984
    end
    local.get 2
    i32.const 1049252
    i32.const 1
    call 15
    drop
    local.get 2
    local.get 7
    local.get 5
    call 15
    drop
    local.get 2
    i32.const 1049257
    i32.const 5
    call 15
    drop
    call 50
    local.set 3
    call 3
    local.set 4
    local.get 1
    local.get 3
    local.get 6
    local.get 0
    i32.load offset=12
    local.get 4
    local.get 2
    local.get 0
    i32.const 40036
    i32.add
    call 53
    local.set 14
    call 68
    call 154
    call 136
    local.get 14
    call 50
    call 65
    local.get 0
    i32.const 80080
    i32.add
    global.set 0)
  (func (;160;) (type 26) (param i32 i32 i32 i32 i32) (result i32)
    (local i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 5
    global.set 0
    block  ;; label = @1
      local.get 0
      local.get 1
      local.get 2
      local.get 3
      local.get 4
      call 192
      local.tee 6
      br_if 0 (;@1;)
      local.get 5
      i32.const 8
      i32.add
      local.get 3
      local.get 0
      local.get 1
      local.get 4
      i32.load offset=12
      call_indirect (type 0)
      i32.const 0
      local.set 6
      local.get 5
      i32.load offset=8
      br_if 0 (;@1;)
      local.get 5
      i32.load offset=12
      local.tee 6
      local.get 2
      i32.load
      i32.store offset=8
      local.get 2
      local.get 6
      i32.store
      local.get 0
      local.get 1
      local.get 2
      local.get 3
      local.get 4
      call 192
      local.set 6
    end
    local.get 5
    i32.const 16
    i32.add
    global.set 0
    local.get 6)
  (func (;161;) (type 8) (param i32) (result i32)
    i32.const 1)
  (func (;162;) (type 8) (param i32) (result i32)
    i32.const 0)
  (func (;163;) (type 14)
    call 164
    unreachable)
  (func (;164;) (type 14)
    call 180
    unreachable)
  (func (;165;) (type 14)
    call 34
    i32.const 0
    call 83)
  (func (;166;) (type 14)
    (local i32 i32 i32 i32 i32 i32 i32 i64)
    global.get 0
    i32.const 112
    i32.sub
    local.tee 0
    global.set 0
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  call 98
                  local.tee 1
                  call 137
                  local.tee 2
                  call 66
                  br_if 0 (;@7;)
                  local.get 2
                  call 51
                  local.set 2
                  local.get 0
                  i32.const 25
                  i32.store offset=84
                  local.get 0
                  i32.const 1048590
                  i32.store offset=80
                  local.get 0
                  i32.const 24
                  i32.add
                  local.get 2
                  call 120
                  local.get 0
                  i32.const 24
                  i32.add
                  call 150
                  local.set 3
                  local.get 0
                  i32.const 24
                  i32.add
                  call 146
                  local.set 2
                  call 3
                  local.set 4
                  block  ;; label = @8
                    loop  ;; label = @9
                      local.get 2
                      i32.eqz
                      br_if 1 (;@8;)
                      local.get 4
                      local.get 0
                      i32.const 24
                      i32.add
                      call 150
                      call 60
                      local.get 2
                      i32.const -1
                      i32.add
                      local.set 2
                      br 0 (;@9;)
                    end
                  end
                  local.get 0
                  i32.load offset=40
                  local.get 0
                  i32.load offset=36
                  i32.ne
                  br_if 1 (;@6;)
                  block  ;; label = @8
                    local.get 0
                    i32.load8_u offset=32
                    i32.eqz
                    br_if 0 (;@8;)
                    i32.const 0
                    i32.const 0
                    i32.store offset=1049944
                    i32.const 0
                    i32.const 0
                    i32.store8 offset=1059948
                  end
                  local.get 1
                  call 3
                  call 26
                  drop
                  local.get 0
                  call 19
                  i32.store offset=20
                  local.get 0
                  i32.const 0
                  i32.store offset=16
                  local.get 0
                  i32.const 104
                  i32.add
                  local.tee 1
                  i64.const 0
                  i64.store
                  local.get 0
                  i32.const 96
                  i32.add
                  local.tee 5
                  i64.const 0
                  i64.store
                  local.get 0
                  i32.const 88
                  i32.add
                  local.tee 6
                  i64.const 0
                  i64.store
                  local.get 0
                  i64.const 0
                  i64.store offset=80
                  local.get 0
                  i32.const 8
                  i32.add
                  local.get 0
                  i32.const 80
                  i32.add
                  i32.const 32
                  local.get 3
                  call 8
                  local.tee 2
                  call 157
                  local.get 3
                  i32.const 0
                  local.get 0
                  i32.load offset=8
                  local.get 0
                  i32.load offset=12
                  call 62
                  drop
                  local.get 0
                  i32.const 52
                  i32.add
                  local.get 1
                  i64.load
                  i64.store align=4
                  local.get 0
                  i32.const 44
                  i32.add
                  local.get 5
                  i64.load
                  i64.store align=4
                  local.get 0
                  i32.const 36
                  i32.add
                  local.get 6
                  i64.load
                  i64.store align=4
                  local.get 0
                  local.get 2
                  i32.store offset=24
                  local.get 0
                  local.get 0
                  i64.load offset=80
                  i64.store offset=28 align=4
                  local.get 2
                  i32.eqz
                  br_if 0 (;@7;)
                  local.get 2
                  i32.const 14
                  i32.ne
                  br_if 6 (;@1;)
                  local.get 0
                  i32.const 24
                  i32.add
                  i32.const 4
                  i32.or
                  i32.const 14
                  i32.const 1049144
                  i32.const 14
                  call 153
                  i32.eqz
                  br_if 6 (;@1;)
                  local.get 0
                  i32.const 64
                  i32.add
                  local.get 4
                  call 81
                  local.get 0
                  i32.const 16
                  i32.add
                  call 152
                  call 18
                  local.tee 7
                  i64.const 4294967296
                  i64.ge_u
                  br_if 2 (;@5;)
                  block  ;; label = @8
                    block  ;; label = @9
                      local.get 7
                      i32.wrap_i64
                      local.tee 4
                      br_if 0 (;@9;)
                      i32.const 1
                      local.set 2
                      local.get 0
                      i32.const 16
                      i32.add
                      call 152
                      call 75
                      call 138
                      local.set 4
                      br 1 (;@8;)
                    end
                    block  ;; label = @9
                      block  ;; label = @10
                        local.get 0
                        i32.load offset=16
                        local.get 0
                        i32.load offset=20
                        i32.lt_s
                        br_if 0 (;@10;)
                        call 3
                        drop
                        br 1 (;@9;)
                      end
                      local.get 0
                      i32.const 16
                      i32.add
                      call 152
                      call 75
                      drop
                    end
                    i32.const 0
                    local.set 2
                  end
                  local.get 0
                  i32.const 0
                  i32.store offset=84
                  local.get 0
                  i32.const 1049856
                  i32.store offset=80
                  local.get 0
                  i32.load offset=72
                  local.get 0
                  i32.load offset=68
                  i32.lt_u
                  br_if 3 (;@4;)
                  local.get 0
                  i32.const 0
                  i32.store offset=84
                  local.get 0
                  i32.const 1049856
                  i32.store offset=80
                  local.get 0
                  i32.load offset=16
                  local.get 0
                  i32.load offset=20
                  i32.lt_s
                  br_if 4 (;@3;)
                  block  ;; label = @8
                    local.get 2
                    i32.eqz
                    br_if 0 (;@8;)
                    call 154
                    local.get 4
                    call 145
                    br 1 (;@7;)
                  end
                  call 70
                  local.set 1
                  block  ;; label = @8
                    block  ;; label = @9
                      block  ;; label = @10
                        block  ;; label = @11
                          call 13
                          br_if 0 (;@11;)
                          call 73
                          local.set 4
                          br 1 (;@10;)
                        end
                        i64.const 0
                        call 2
                        local.tee 4
                        call 35
                        local.get 0
                        i32.const 104
                        i32.add
                        i64.const 0
                        i64.store
                        local.get 0
                        i32.const 96
                        i32.add
                        i64.const 0
                        i64.store
                        local.get 0
                        i32.const 88
                        i32.add
                        i64.const 0
                        i64.store
                        local.get 0
                        i64.const 0
                        i64.store offset=80
                        local.get 0
                        i32.const 80
                        i32.add
                        call 36
                        local.tee 2
                        br_if 1 (;@9;)
                      end
                      call 3
                      local.set 2
                      br 1 (;@8;)
                    end
                    local.get 2
                    i32.const 33
                    i32.ge_u
                    br_if 6 (;@2;)
                    local.get 0
                    i32.const 80
                    i32.add
                    local.get 2
                    call 7
                    local.set 2
                  end
                  local.get 2
                  call 66
                  i32.eqz
                  br_if 0 (;@7;)
                  local.get 4
                  i64.const 0
                  call 123
                  i32.const 255
                  i32.and
                  i32.const 1
                  i32.ne
                  br_if 0 (;@7;)
                  local.get 1
                  local.get 2
                  i64.const 0
                  local.get 4
                  call 65
                end
                local.get 0
                i32.const 112
                i32.add
                global.set 0
                return
              end
              local.get 0
              i32.const 80
              i32.add
              i32.const 1048576
              i32.const 14
              call 147
              unreachable
            end
            i32.const 1049166
            i32.const 6
            i32.const 1048576
            i32.const 14
            call 76
            unreachable
          end
          local.get 0
          i32.const 80
          i32.add
          i32.const 1049029
          i32.const 18
          call 147
          unreachable
        end
        local.get 0
        i32.const 80
        i32.add
        i32.const 1049029
        i32.const 18
        call 147
        unreachable
      end
      local.get 2
      i32.const 32
      call 48
      unreachable
    end
    i32.const 1049198
    i32.const 54
    call 1
    unreachable)
  (func (;167;) (type 14)
    (local i32 i32 i64)
    global.get 0
    i32.const 40032
    i32.sub
    local.tee 0
    global.set 0
    call 74
    i32.const 0
    call 83
    call 73
    local.set 1
    local.get 0
    call 155
    call 139
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 1
          local.get 0
          i32.load offset=8
          call 121
          i32.eqz
          br_if 0 (;@3;)
          local.get 0
          i32.const 40016
          i32.add
          i32.load
          i32.eqz
          br_if 1 (;@2;)
          block  ;; label = @4
            local.get 0
            i64.load
            local.tee 2
            i64.eqz
            br_if 0 (;@4;)
            local.get 2
            call 37
            i64.gt_u
            br_if 3 (;@1;)
          end
          call 159
          local.get 0
          i32.const 40032
          i32.add
          global.set 0
          return
        end
        i32.const 1049618
        i32.const 36
        call 52
        unreachable
      end
      i32.const 1049654
      i32.const 22
      call 52
      unreachable
    end
    i32.const 1049469
    i32.const 22
    call 52
    unreachable)
  (func (;168;) (type 14)
    (local i32 i32 i32 i32 i64)
    global.get 0
    i32.const 40064
    i32.sub
    local.tee 0
    global.set 0
    call 74
    i32.const 1
    call 83
    i32.const 0
    i32.const 1049158
    i32.const 8
    call 79
    local.set 1
    call 73
    local.set 2
    local.get 0
    i32.const 16
    i32.add
    call 155
    call 139
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            local.get 1
            i32.eqz
            br_if 0 (;@4;)
            local.get 0
            i32.const 40032
            i32.add
            i32.load
            local.get 1
            i32.lt_u
            br_if 1 (;@3;)
            local.get 1
            i64.extend_i32_u
            call 2
            local.tee 3
            local.get 3
            local.get 0
            i32.load offset=24
            call 44
            call 38
            local.get 3
            local.get 2
            call 121
            i32.eqz
            br_if 2 (;@2;)
            block  ;; label = @5
              local.get 0
              i64.load offset=16
              local.tee 4
              i64.eqz
              br_if 0 (;@5;)
              local.get 4
              call 37
              i64.gt_u
              br_if 4 (;@1;)
            end
            local.get 0
            i32.const 0
            i32.store8 offset=40056
            local.get 0
            i32.const 0
            i32.store offset=40048
            local.get 0
            local.get 1
            i32.const -1
            i32.add
            i32.store offset=40052
            block  ;; label = @5
              loop  ;; label = @6
                local.get 0
                i32.const 8
                i32.add
                local.get 0
                i32.const 40048
                i32.add
                call 156
                local.get 0
                i32.load offset=8
                i32.eqz
                br_if 1 (;@5;)
                call 159
                br 0 (;@6;)
              end
            end
            local.get 0
            i32.const 40064
            i32.add
            global.set 0
            return
          end
          i32.const 1049320
          i32.const 40
          call 52
          unreachable
        end
        i32.const 1049360
        i32.const 24
        call 52
        unreachable
      end
      i32.const 1049384
      i32.const 85
      call 52
      unreachable
    end
    i32.const 1049469
    i32.const 22
    call 52
    unreachable)
  (func (;169;) (type 14)
    (local i32 i32 i32)
    call 34
    call 71
    i32.const 0
    call 83
    call 68
    local.set 0
    call 63
    local.set 1
    i64.const 0
    call 2
    local.set 2
    local.get 1
    i32.const 1059952
    call 33
    drop
    i32.const 1059952
    local.get 2
    call 39
    local.get 0
    local.get 2
    i64.const 0
    i32.const 1049315
    i32.const 5
    call 7
    call 3
    call 9
    drop)
  (func (;170;) (type 14)
    (local i32 i32 i32 i64 i32 i32 i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 0
    global.set 0
    call 34
    i32.const 1
    call 83
    i32.const 0
    i32.const 1049158
    i32.const 8
    call 79
    local.set 1
    local.get 0
    call 3
    i32.store offset=12
    call 154
    call 136
    local.set 2
    local.get 0
    i32.const 0
    i32.store8 offset=24
    local.get 0
    local.get 1
    i32.const -1
    i32.add
    i32.store offset=20
    local.get 0
    i32.const 0
    i32.store offset=16
    block  ;; label = @1
      loop  ;; label = @2
        local.get 0
        local.get 0
        i32.const 16
        i32.add
        call 156
        local.get 0
        i32.load
        i32.eqz
        br_if 1 (;@1;)
        local.get 2
        call 50
        call 3
        call 50
        call 3
        call 3
        local.get 0
        i32.const 12
        i32.add
        call 53
        local.set 3
        call 50
        local.set 4
        call 3
        local.tee 1
        local.get 2
        call 54
        block  ;; label = @3
          block  ;; label = @4
            local.get 3
            i64.eqz
            i32.eqz
            br_if 0 (;@4;)
            i32.const 1048626
            local.set 5
            i32.const 13
            local.set 6
            br 1 (;@3;)
          end
          local.get 1
          local.get 3
          call 114
          i32.const 1048615
          local.set 5
          i32.const 11
          local.set 6
        end
        local.get 1
        local.get 4
        call 55
        call 6
        local.get 5
        local.get 6
        call 7
        local.get 1
        call 61
        drop
        br 0 (;@2;)
      end
    end
    local.get 0
    i32.const 32
    i32.add
    global.set 0)
  (func (;171;) (type 14)
    (local i32 i32 i32 i32 i64)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 0
    global.set 0
    call 34
    i32.const 2
    call 83
    i32.const 0
    i32.const 1049158
    i32.const 8
    call 79
    local.set 1
    i32.const 1
    call 80
    local.set 2
    local.get 0
    call 3
    i32.store offset=12
    call 154
    call 136
    local.set 3
    local.get 0
    i32.const 0
    i32.store8 offset=24
    local.get 0
    local.get 1
    i32.const -1
    i32.add
    i32.store offset=20
    local.get 0
    i32.const 0
    i32.store offset=16
    block  ;; label = @1
      loop  ;; label = @2
        local.get 0
        local.get 0
        i32.const 16
        i32.add
        call 156
        local.get 0
        i32.load
        i32.eqz
        br_if 1 (;@1;)
        local.get 3
        call 50
        call 3
        call 50
        call 3
        call 3
        local.get 0
        i32.const 12
        i32.add
        call 53
        local.set 4
        local.get 2
        call 154
        call 136
        local.get 4
        call 50
        call 65
        br 0 (;@2;)
      end
    end
    local.get 0
    i32.const 32
    i32.add
    global.set 0)
  (func (;172;) (type 14)
    (local i32 i32 i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 0
    global.set 0
    call 34
    i32.const 0
    local.set 1
    i32.const 0
    call 83
    local.get 0
    call 3
    i32.store offset=12
    call 154
    call 136
    local.set 2
    i32.const 0
    local.set 3
    block  ;; label = @1
      loop  ;; label = @2
        local.get 3
        i32.const 255
        i32.and
        br_if 1 (;@1;)
        local.get 2
        call 50
        call 3
        call 50
        call 3
        call 3
        local.get 0
        i32.const 12
        i32.add
        call 53
        drop
        i32.const 99
        local.get 1
        i32.const 1
        i32.add
        local.get 1
        i32.const 99
        i32.eq
        local.tee 3
        select
        local.set 1
        br 0 (;@2;)
      end
    end
    local.get 0
    i32.const 16
    i32.add
    global.set 0)
  (func (;173;) (type 14)
    (local i32 i32 i64 i64)
    global.get 0
    i32.const 40032
    i32.sub
    local.tee 0
    global.set 0
    call 34
    i32.const 0
    call 83
    call 155
    call 3
    local.tee 1
    call 25
    drop
    block  ;; label = @1
      local.get 1
      call 8
      i32.eqz
      br_if 0 (;@1;)
      local.get 0
      call 155
      call 139
      local.get 0
      i32.const 40016
      i32.add
      i64.load32_u
      local.set 2
      local.get 0
      i64.load
      local.set 3
      local.get 0
      i32.load offset=8
      call 40
      local.get 2
      call 41
      local.get 3
      call 41
    end
    local.get 0
    i32.const 40032
    i32.add
    global.set 0)
  (func (;174;) (type 14)
    (local i32 i32 i32 i32)
    global.get 0
    i32.const 80
    i32.sub
    local.tee 0
    global.set 0
    call 74
    call 71
    i32.const 2
    call 83
    i32.const 0
    call 75
    local.set 1
    i32.const 1
    call 75
    local.set 2
    call 72
    local.set 3
    local.get 0
    i32.const 256
    i32.store16 offset=20
    local.get 0
    i32.const 65793
    i32.store offset=16
    local.get 0
    i32.const 48
    i32.add
    local.get 3
    local.get 1
    local.get 2
    local.get 0
    i32.const 16
    i32.add
    call 90
    local.get 0
    i32.const 16
    i32.add
    local.get 0
    i32.const 48
    i32.add
    call 104
    local.get 0
    i32.const 8
    i32.add
    call 158
    local.get 0
    i32.const 68
    i32.add
    local.get 0
    i64.load offset=8
    i64.store align=4
    local.get 0
    i32.const 1
    i32.store offset=64
    local.get 0
    local.get 0
    i64.load offset=16
    i64.store offset=48
    local.get 0
    local.get 0
    i64.load offset=24
    i64.store offset=56
    local.get 0
    i32.const 48
    i32.add
    call 97
    unreachable)
  (func (;175;) (type 14)
    (local i32 i32 i32 i32 i32 i32 i32 i64 i32)
    global.get 0
    i32.const 80064
    i32.sub
    local.tee 0
    global.set 0
    call 34
    call 71
    i32.const 7
    call 83
    i32.const 0
    call 78
    local.set 1
    i32.const 1
    call 78
    local.set 2
    i32.const 2
    i32.const 1049172
    i32.const 22
    call 79
    local.set 3
    i32.const 3
    call 75
    local.set 4
    i32.const 4
    call 75
    local.set 5
    i32.const 5
    call 75
    local.set 6
    i32.const 6
    call 18
    local.set 7
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              local.get 1
              call 127
              i32.eqz
              br_if 0 (;@5;)
              local.get 2
              i64.const 10000
              call 123
              i32.const 255
              i32.and
              i32.const 1
              i32.eq
              br_if 1 (;@4;)
              local.get 4
              call 66
              br_if 2 (;@3;)
              local.get 6
              call 66
              br_if 3 (;@2;)
              local.get 0
              i32.const 0
              i32.store offset=40016
              local.get 0
              i32.const 0
              i32.store8 offset=40032
              local.get 0
              local.get 3
              i32.const -1
              i32.add
              i32.store offset=40028
              local.get 0
              i32.const 0
              i32.store offset=40024
              local.get 0
              i32.const 16
              i32.add
              local.set 8
              i32.const 1
              local.set 3
              block  ;; label = @6
                loop  ;; label = @7
                  local.get 0
                  i32.const 8
                  i32.add
                  local.get 0
                  i32.const 40024
                  i32.add
                  call 156
                  local.get 0
                  i32.load offset=8
                  i32.eqz
                  br_if 1 (;@6;)
                  local.get 3
                  i32.const 10001
                  i32.eq
                  br_if 6 (;@1;)
                  local.get 8
                  local.get 0
                  i32.load offset=12
                  i32.store
                  local.get 8
                  i32.const 4
                  i32.add
                  local.set 8
                  local.get 0
                  local.get 3
                  i32.store offset=40016
                  local.get 3
                  i32.const 1
                  i32.add
                  local.set 3
                  br 0 (;@7;)
                end
              end
              local.get 0
              i32.const 40040
              i32.add
              local.get 0
              i32.const 16
              i32.add
              i32.const 40004
              call 200
              drop
              local.get 0
              local.get 6
              i32.store offset=80052
              local.get 0
              local.get 5
              i32.store offset=80048
              local.get 0
              local.get 4
              i32.store offset=80044
              local.get 0
              local.get 2
              i32.store offset=40036
              local.get 0
              local.get 1
              i32.store offset=40032
              local.get 0
              local.get 7
              i64.store offset=40024
              call 155
              local.get 0
              i32.const 40024
              i32.add
              call 142
              local.get 0
              i32.const 80064
              i32.add
              global.set 0
              return
            end
            i32.const 1049493
            i32.const 25
            call 52
            unreachable
          end
          i32.const 1049518
          i32.const 44
          call 52
          unreachable
        end
        i32.const 1049562
        i32.const 25
        call 52
        unreachable
      end
      i32.const 1049587
      i32.const 31
      call 52
      unreachable
    end
    local.get 0
    i32.const 80060
    i32.add
    call 176
    unreachable)
  (func (;176;) (type 10) (param i32)
    call 164
    unreachable)
  (func (;177;) (type 14)
    (local i32)
    global.get 0
    i32.const 64
    i32.sub
    local.tee 0
    global.set 0
    call 34
    call 71
    i32.const 0
    call 83
    local.get 0
    i32.const 32
    i32.add
    call 69
    call 154
    call 136
    call 86
    local.get 0
    local.get 0
    i32.const 32
    i32.add
    call 104
    local.get 0
    call 97
    unreachable)
  (func (;178;) (type 14)
    (local i32)
    call 34
    i32.const 1
    call 83
    i32.const 0
    call 75
    local.set 0
    call 154
    local.get 0
    call 145)
  (func (;179;) (type 14)
    (local i32 i32 i32 i32 i32 i64)
    global.get 0
    i32.const 80
    i32.sub
    local.tee 0
    global.set 0
    call 34
    i32.const 1
    call 83
    i32.const 0
    local.set 1
    i32.const 0
    call 80
    local.set 2
    local.get 0
    call 3
    i32.store offset=12
    call 154
    call 136
    local.set 3
    i32.const 0
    local.set 4
    block  ;; label = @1
      loop  ;; label = @2
        local.get 4
        i32.const 255
        i32.and
        br_if 1 (;@1;)
        local.get 3
        call 50
        call 3
        call 50
        call 3
        call 3
        local.get 0
        i32.const 12
        i32.add
        call 53
        drop
        i32.const 99
        local.get 1
        i32.const 1
        i32.add
        local.get 1
        i32.const 99
        i32.eq
        local.tee 4
        select
        local.set 1
        br 0 (;@2;)
      end
    end
    i32.const 1048976
    call 96
    drop
    call 3
    local.set 1
    local.get 0
    i32.const 16
    i32.add
    local.get 2
    i32.const 1049127
    i32.const 17
    call 7
    local.get 1
    call 106
    local.get 0
    i32.const 48
    i32.add
    i32.const 24
    i32.add
    local.get 0
    i32.const 16
    i32.add
    i32.const 24
    i32.add
    i64.load
    i64.store
    local.get 0
    i32.const 48
    i32.add
    i32.const 16
    i32.add
    local.get 0
    i32.const 16
    i32.add
    i32.const 16
    i32.add
    i64.load
    i64.store
    local.get 0
    i32.const 48
    i32.add
    i32.const 8
    i32.add
    local.get 0
    i32.const 16
    i32.add
    i32.const 8
    i32.add
    i64.load
    i64.store
    local.get 0
    local.get 0
    i64.load offset=16
    i64.store offset=48
    local.get 0
    i32.const 16
    i32.add
    local.get 0
    i32.const 48
    i32.add
    call 105
    block  ;; label = @1
      local.get 0
      i64.load offset=16
      local.tee 5
      i64.const -1
      i64.ne
      br_if 0 (;@1;)
      call 6
      local.set 5
    end
    local.get 0
    i32.const 48
    i32.add
    local.get 5
    local.get 0
    i32.load offset=24
    local.get 0
    i32.load offset=28
    local.get 0
    i32.load offset=36
    local.get 0
    i32.load offset=40
    call 64
    call 81
    local.get 0
    i32.const 80
    i32.add
    global.set 0)
  (func (;180;) (type 14)
    i32.const 1049817
    i32.const 14
    call 1
    unreachable)
  (func (;181;) (type 10) (param i32)
    local.get 0
    call 182
    unreachable)
  (func (;182;) (type 10) (param i32)
    local.get 0
    call 183
    unreachable)
  (func (;183;) (type 10) (param i32)
    call 164
    unreachable)
  (func (;184;) (type 3) (param i32 i32)
    local.get 0
    local.get 1
    call 185
    unreachable)
  (func (;185;) (type 3) (param i32 i32)
    local.get 0
    local.get 1
    call 186
    unreachable)
  (func (;186;) (type 3) (param i32 i32)
    call 164
    unreachable)
  (func (;187;) (type 3) (param i32 i32)
    local.get 0
    local.get 1
    call 188
    unreachable)
  (func (;188;) (type 3) (param i32 i32)
    local.get 0
    local.get 1
    call 189
    unreachable)
  (func (;189;) (type 3) (param i32 i32)
    call 164
    unreachable)
  (func (;190;) (type 3) (param i32 i32)
    call 164
    unreachable)
  (func (;191;) (type 0) (param i32 i32 i32 i32)
    (local i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 4
    global.set 0
    local.get 4
    local.get 1
    i32.load
    local.tee 1
    i32.load
    i32.store offset=12
    local.get 2
    i32.const 2
    i32.add
    local.tee 2
    local.get 2
    i32.mul
    local.tee 2
    i32.const 2048
    local.get 2
    i32.const 2048
    i32.gt_u
    select
    local.tee 5
    i32.const 4
    local.get 4
    i32.const 12
    i32.add
    i32.const 1049856
    i32.const 1049856
    call 160
    local.set 2
    local.get 1
    local.get 4
    i32.load offset=12
    i32.store
    block  ;; label = @1
      block  ;; label = @2
        local.get 2
        br_if 0 (;@2;)
        i32.const 1
        local.set 1
        br 1 (;@1;)
      end
      local.get 2
      i64.const 0
      i64.store offset=4 align=4
      local.get 2
      local.get 2
      local.get 5
      i32.const 2
      i32.shl
      i32.add
      i32.const 2
      i32.or
      i32.store
      i32.const 0
      local.set 1
    end
    local.get 0
    local.get 2
    i32.store offset=4
    local.get 0
    local.get 1
    i32.store
    local.get 4
    i32.const 16
    i32.add
    global.set 0)
  (func (;192;) (type 26) (param i32 i32 i32 i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32)
    local.get 1
    i32.const -1
    i32.add
    local.set 5
    i32.const 0
    local.set 6
    i32.const 0
    local.get 1
    i32.sub
    local.set 7
    local.get 0
    i32.const 2
    i32.shl
    local.set 8
    local.get 2
    i32.load
    local.set 9
    block  ;; label = @1
      loop  ;; label = @2
        local.get 9
        i32.eqz
        br_if 1 (;@1;)
        local.get 9
        local.set 1
        loop  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                local.get 1
                i32.load offset=8
                local.tee 9
                i32.const 1
                i32.and
                br_if 0 (;@6;)
                local.get 1
                i32.load
                i32.const -4
                i32.and
                local.tee 10
                local.get 1
                i32.const 8
                i32.add
                local.tee 11
                i32.sub
                local.get 8
                i32.lt_u
                br_if 1 (;@5;)
                block  ;; label = @7
                  block  ;; label = @8
                    local.get 11
                    local.get 3
                    local.get 0
                    local.get 4
                    i32.load offset=16
                    call_indirect (type 1)
                    i32.const 2
                    i32.shl
                    i32.add
                    i32.const 8
                    i32.add
                    local.get 10
                    local.get 8
                    i32.sub
                    local.get 7
                    i32.and
                    local.tee 9
                    i32.le_u
                    br_if 0 (;@8;)
                    local.get 11
                    i32.load
                    local.set 9
                    local.get 5
                    local.get 11
                    i32.and
                    br_if 3 (;@5;)
                    local.get 2
                    local.get 9
                    i32.const -4
                    i32.and
                    i32.store
                    local.get 1
                    local.get 1
                    i32.load
                    i32.const 1
                    i32.or
                    i32.store
                    local.get 1
                    local.set 9
                    br 1 (;@7;)
                  end
                  local.get 9
                  i32.const 0
                  i32.store
                  local.get 9
                  i32.const -8
                  i32.add
                  local.tee 9
                  i64.const 0
                  i64.store align=4
                  local.get 9
                  local.get 1
                  i32.load
                  i32.const -4
                  i32.and
                  i32.store
                  block  ;; label = @8
                    local.get 1
                    i32.load
                    local.tee 2
                    i32.const -4
                    i32.and
                    local.tee 11
                    i32.eqz
                    br_if 0 (;@8;)
                    i32.const 0
                    local.get 11
                    local.get 2
                    i32.const 2
                    i32.and
                    select
                    local.tee 2
                    i32.eqz
                    br_if 0 (;@8;)
                    local.get 2
                    local.get 2
                    i32.load offset=4
                    i32.const 3
                    i32.and
                    local.get 9
                    i32.or
                    i32.store offset=4
                  end
                  local.get 9
                  local.get 9
                  i32.load offset=4
                  i32.const 3
                  i32.and
                  local.get 1
                  i32.or
                  i32.store offset=4
                  local.get 1
                  local.get 1
                  i32.load offset=8
                  i32.const -2
                  i32.and
                  i32.store offset=8
                  local.get 1
                  local.get 1
                  i32.load
                  local.tee 2
                  i32.const 3
                  i32.and
                  local.get 9
                  i32.or
                  local.tee 11
                  i32.store
                  block  ;; label = @8
                    block  ;; label = @9
                      local.get 2
                      i32.const 2
                      i32.and
                      br_if 0 (;@9;)
                      local.get 9
                      i32.load
                      local.set 1
                      br 1 (;@8;)
                    end
                    local.get 1
                    local.get 11
                    i32.const -3
                    i32.and
                    i32.store
                    local.get 9
                    local.get 9
                    i32.load
                    i32.const 2
                    i32.or
                    local.tee 1
                    i32.store
                  end
                  local.get 9
                  local.get 1
                  i32.const 1
                  i32.or
                  i32.store
                end
                local.get 9
                i32.const 8
                i32.add
                local.set 6
                br 5 (;@1;)
              end
              local.get 1
              local.get 9
              i32.const -2
              i32.and
              i32.store offset=8
              block  ;; label = @6
                block  ;; label = @7
                  local.get 1
                  i32.load offset=4
                  i32.const -4
                  i32.and
                  local.tee 9
                  br_if 0 (;@7;)
                  i32.const 0
                  local.set 9
                  br 1 (;@6;)
                end
                i32.const 0
                local.get 9
                local.get 9
                i32.load8_u
                i32.const 1
                i32.and
                select
                local.set 9
              end
              local.get 1
              call 193
              local.get 1
              i32.load8_u
              i32.const 2
              i32.and
              i32.eqz
              br_if 1 (;@4;)
              local.get 9
              local.get 9
              i32.load
              i32.const 2
              i32.or
              i32.store
              br 1 (;@4;)
            end
            local.get 2
            local.get 9
            i32.store
            br 2 (;@2;)
          end
          local.get 2
          local.get 9
          i32.store
          local.get 9
          local.set 1
          br 0 (;@3;)
        end
      end
    end
    local.get 6)
  (func (;193;) (type 10) (param i32)
    (local i32 i32)
    block  ;; label = @1
      local.get 0
      i32.load
      local.tee 1
      i32.const -4
      i32.and
      local.tee 2
      i32.eqz
      br_if 0 (;@1;)
      i32.const 0
      local.get 2
      local.get 1
      i32.const 2
      i32.and
      select
      local.tee 1
      i32.eqz
      br_if 0 (;@1;)
      local.get 1
      local.get 1
      i32.load offset=4
      i32.const 3
      i32.and
      local.get 0
      i32.load offset=4
      i32.const -4
      i32.and
      i32.or
      i32.store offset=4
    end
    block  ;; label = @1
      local.get 0
      i32.load offset=4
      local.tee 1
      i32.const -4
      i32.and
      local.tee 2
      i32.eqz
      br_if 0 (;@1;)
      local.get 2
      local.get 2
      i32.load
      i32.const 3
      i32.and
      local.get 0
      i32.load
      i32.const -4
      i32.and
      i32.or
      i32.store
      local.get 0
      i32.load offset=4
      local.set 1
    end
    local.get 0
    local.get 1
    i32.const 3
    i32.and
    i32.store offset=4
    local.get 0
    local.get 0
    i32.load
    i32.const 3
    i32.and
    i32.store)
  (func (;194;) (type 10) (param i32))
  (func (;195;) (type 0) (param i32 i32 i32 i32)
    (local i32)
    block  ;; label = @1
      block  ;; label = @2
        local.get 2
        i32.const 2
        i32.shl
        local.tee 2
        local.get 3
        i32.const 3
        i32.shl
        i32.const 16384
        i32.add
        local.tee 3
        local.get 2
        local.get 3
        i32.gt_u
        select
        i32.const 65543
        i32.add
        local.tee 4
        i32.const 16
        i32.shr_u
        memory.grow
        local.tee 3
        i32.const -1
        i32.ne
        br_if 0 (;@2;)
        i32.const 1
        local.set 2
        br 1 (;@1;)
      end
      local.get 3
      i32.const 16
      i32.shl
      local.tee 3
      i64.const 0
      i64.store
      i32.const 0
      local.set 2
      local.get 3
      i32.const 0
      i32.store offset=8
      local.get 3
      local.get 3
      local.get 4
      i32.const -65536
      i32.and
      i32.add
      i32.const 2
      i32.or
      i32.store
    end
    local.get 0
    local.get 3
    i32.store offset=4
    local.get 0
    local.get 2
    i32.store)
  (func (;196;) (type 1) (param i32 i32) (result i32)
    i32.const 512)
  (func (;197;) (type 1) (param i32 i32) (result i32)
    local.get 1)
  (func (;198;) (type 10) (param i32))
  (func (;199;) (type 6) (param i32 i32 i32) (result i32)
    local.get 0
    local.get 1
    local.get 2
    call 202)
  (func (;200;) (type 6) (param i32 i32 i32) (result i32)
    local.get 0
    local.get 1
    local.get 2
    call 201)
  (func (;201;) (type 6) (param i32 i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32)
    block  ;; label = @1
      block  ;; label = @2
        local.get 2
        i32.const 15
        i32.gt_u
        br_if 0 (;@2;)
        local.get 0
        local.set 3
        br 1 (;@1;)
      end
      local.get 0
      i32.const 0
      local.get 0
      i32.sub
      i32.const 3
      i32.and
      local.tee 4
      i32.add
      local.set 5
      block  ;; label = @2
        local.get 4
        i32.eqz
        br_if 0 (;@2;)
        local.get 0
        local.set 3
        local.get 1
        local.set 6
        loop  ;; label = @3
          local.get 3
          local.get 6
          i32.load8_u
          i32.store8
          local.get 6
          i32.const 1
          i32.add
          local.set 6
          local.get 3
          i32.const 1
          i32.add
          local.tee 3
          local.get 5
          i32.lt_u
          br_if 0 (;@3;)
        end
      end
      local.get 5
      local.get 2
      local.get 4
      i32.sub
      local.tee 7
      i32.const -4
      i32.and
      local.tee 8
      i32.add
      local.set 3
      block  ;; label = @2
        block  ;; label = @3
          local.get 1
          local.get 4
          i32.add
          local.tee 9
          i32.const 3
          i32.and
          i32.eqz
          br_if 0 (;@3;)
          local.get 8
          i32.const 1
          i32.lt_s
          br_if 1 (;@2;)
          local.get 9
          i32.const 3
          i32.shl
          local.tee 6
          i32.const 24
          i32.and
          local.set 2
          local.get 9
          i32.const -4
          i32.and
          local.tee 10
          i32.const 4
          i32.add
          local.set 1
          i32.const 0
          local.get 6
          i32.sub
          i32.const 24
          i32.and
          local.set 4
          local.get 10
          i32.load
          local.set 6
          loop  ;; label = @4
            local.get 5
            local.get 6
            local.get 2
            i32.shr_u
            local.get 1
            i32.load
            local.tee 6
            local.get 4
            i32.shl
            i32.or
            i32.store
            local.get 1
            i32.const 4
            i32.add
            local.set 1
            local.get 5
            i32.const 4
            i32.add
            local.tee 5
            local.get 3
            i32.lt_u
            br_if 0 (;@4;)
            br 2 (;@2;)
          end
        end
        local.get 8
        i32.const 1
        i32.lt_s
        br_if 0 (;@2;)
        local.get 9
        local.set 1
        loop  ;; label = @3
          local.get 5
          local.get 1
          i32.load
          i32.store
          local.get 1
          i32.const 4
          i32.add
          local.set 1
          local.get 5
          i32.const 4
          i32.add
          local.tee 5
          local.get 3
          i32.lt_u
          br_if 0 (;@3;)
        end
      end
      local.get 7
      i32.const 3
      i32.and
      local.set 2
      local.get 9
      local.get 8
      i32.add
      local.set 1
    end
    block  ;; label = @1
      local.get 2
      i32.eqz
      br_if 0 (;@1;)
      local.get 3
      local.get 2
      i32.add
      local.set 5
      loop  ;; label = @2
        local.get 3
        local.get 1
        i32.load8_u
        i32.store8
        local.get 1
        i32.const 1
        i32.add
        local.set 1
        local.get 3
        i32.const 1
        i32.add
        local.tee 3
        local.get 5
        i32.lt_u
        br_if 0 (;@2;)
      end
    end
    local.get 0)
  (func (;202;) (type 6) (param i32 i32 i32) (result i32)
    (local i32 i32 i32)
    i32.const 0
    local.set 3
    block  ;; label = @1
      local.get 2
      i32.eqz
      br_if 0 (;@1;)
      block  ;; label = @2
        loop  ;; label = @3
          local.get 0
          i32.load8_u
          local.tee 4
          local.get 1
          i32.load8_u
          local.tee 5
          i32.ne
          br_if 1 (;@2;)
          local.get 0
          i32.const 1
          i32.add
          local.set 0
          local.get 1
          i32.const 1
          i32.add
          local.set 1
          local.get 2
          i32.const -1
          i32.add
          local.tee 2
          i32.eqz
          br_if 2 (;@1;)
          br 0 (;@3;)
        end
      end
      local.get 4
      local.get 5
      i32.sub
      local.set 3
    end
    local.get 3)
  (table (;0;) 9 9 funcref)
  (memory (;0;) 17)
  (global (;0;) (mut i32) (i32.const 1048576))
  (global (;1;) i32 (i32.const 1061012))
  (global (;2;) i32 (i32.const 1061024))
  (export "memory" (memory 0))
  (export "init" (func 165))
  (export "callBack" (func 166))
  (export "buy" (func 167))
  (export "buyMultiple" (func 168))
  (export "claimFunds" (func 169))
  (export "createAndBurn" (func 170))
  (export "createAndSend" (func 171))
  (export "createNoParameter" (func 172))
  (export "getCollectionInfo" (func 173))
  (export "issueToken" (func 174))
  (export "setCollectionInfo" (func 175))
  (export "setLocalRoles" (func 177))
  (export "setTokenID" (func 178))
  (export "tryMore" (func 179))
  (export "__data_end" (global 1))
  (export "__heap_base" (global 2))
  (elem (;0;) (i32.const 1) func 198 191 197 162 194 195 196 161)
  (data (;0;) (i32.const 1048576) "input too longserializer decode error: DCDTNFTBurnDCDTLocalBurnDCDTNFTCreateEndpoint can only be called by ownerargument decode error (): function does not accept DCDT paymentwrong number of argumentssetSpecialRole\00\00\00\00\00\00\00\00\00\01\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\02\ff\ffissueNonFungiblecanFreezecanWipecanPausecanChangeOwnercanUpgradecanAddSpecialRolesMultiDCDTNFTTransferDCDTNFTTransferDCDTTransferCB_CLOSUREinput too short\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00REWAcapacity exceededtoo many argumentsstorage decode error: bad array lengthtoo few argumentsnftTokenIdcollectionsDatacreateNoParameterissue_callbacknum_nftsresultnum_elem_in_collectionaddrno callback function with that name exists in contract/.png.jsoncollection.jsontags:;metadata: # numbat,nft,somethingfundsnumber of nfts must be greater than zeronot enough nfts remainedpayment does not correspond with the number of nfts which is intended to be purchasedbuy is not enabled yet\03\05selling price cannot be 0royalties percent cannot be greater than 100ipfs path cannot be emptynfts start name cannot be emptythe payment must match the fixed sumall the nfts were soldtruefalseDCDTRoleNFTUpdateAttributesDCDTRoleNFTAddURIDCDTRoleNFTBurnDCDTRoleNFTAddQuantityDCDTRoleNFTCreateDCDTRoleLocalBurnDCDTRoleLocalMintpanic occurred\00\01\00\00\00\04\00\00\00\04\00\00\00\02\00\00\00\03\00\00\00\04\00\00\00\05\00\00\00\00\00\00\00\01\00\00\00\06\00\00\00\07\00\00\00\08\00\00\00\00\00\00\00\11\00\00\00\11\00\00\00\11\00\00\00\16\00\00\00\0f\00\00\00\11\00\00\00\1b\00\00\00\00\05\10\00\c8\04\10\00\b7\04\10\00\a6\04\10\00\90\04\10\00\81\04\10\00p\04\10\00U\04\10\00"))
