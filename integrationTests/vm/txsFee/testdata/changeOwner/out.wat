(module
  (type (;0;) (func))
  (type (;1;) (func (param i32)))
  (type (;2;) (func (result i32)))
  (type (;3;) (func (param i32 i32) (result i32)))
  (type (;4;) (func (param i32 i32)))
  (type (;5;) (func (param i32 i32 i32) (result i32)))
  (type (;6;) (func (param i32) (result i32)))
  (type (;7;) (func (param i32 i32 i32 i32)))
  (type (;8;) (func (param i32 i32 i32)))
  (type (;9;) (func (param i32 i64)))
  (type (;10;) (func (param i32 i32 i64 i32 i32) (result i32)))
  (type (;11;) (func (param i32 i32 i32 i32) (result i32)))
  (type (;12;) (func (result i64)))
  (type (;13;) (func (param i64 i32 i32 i32 i32 i32 i32) (result i32)))
  (import "env" "bigIntSetInt64" (func (;0;) (type 9)))
  (import "env" "mBufferNew" (func (;1;) (type 2)))
  (import "env" "mBufferAppend" (func (;2;) (type 3)))
  (import "env" "managedTransferValueExecute" (func (;3;) (type 10)))
  (import "env" "mBufferAppendBytes" (func (;4;) (type 5)))
  (import "env" "managedSignalError" (func (;5;) (type 1)))
  (import "env" "getNumDCDTTransfers" (func (;6;) (type 2)))
  (import "env" "signalError" (func (;7;) (type 4)))
  (import "env" "mBufferGetArgument" (func (;8;) (type 3)))
  (import "env" "mBufferGetLength" (func (;9;) (type 6)))
  (import "env" "getNumArguments" (func (;10;) (type 2)))
  (import "env" "managedGetOriginalTxHash" (func (;11;) (type 1)))
  (import "env" "mBufferGetByteSlice" (func (;12;) (type 11)))
  (import "env" "mBufferStorageStore" (func (;13;) (type 3)))
  (import "env" "mBufferSetBytes" (func (;14;) (type 5)))
  (import "env" "managedAsyncCall" (func (;15;) (type 7)))
  (import "env" "bigIntGetCallValue" (func (;16;) (type 1)))
  (import "env" "checkNoPayment" (func (;17;) (type 0)))
  (import "env" "getGasLeft" (func (;18;) (type 12)))
  (import "env" "mBufferNewFromBytes" (func (;19;) (type 3)))
  (import "env" "managedDeployFromSourceContract" (func (;20;) (type 13)))
  (import "env" "cleanReturnData" (func (;21;) (type 0)))
  (func (;22;) (type 2) (result i32)
    (local i32)
    i32.const 131432
    i32.const 131432
    i32.load
    i32.const 1
    i32.sub
    local.tee 0
    i32.store
    local.get 0)
  (func (;23;) (type 6) (param i32) (result i32)
    (local i32)
    call 1
    local.tee 1
    local.get 0
    call 2
    drop
    local.get 1)
  (func (;24;) (type 4) (param i32 i32)
    local.get 0
    local.get 1
    i64.const 0
    call 25
    call 25
    call 3
    drop)
  (func (;25;) (type 2) (result i32)
    (local i32)
    call 22
    local.tee 0
    i32.const 131408
    i32.const 0
    call 14
    drop
    local.get 0)
  (func (;26;) (type 7) (param i32 i32 i32 i32)
    local.get 2
    local.get 3
    call 28
    local.set 2
    local.get 0
    call 25
    i32.store offset=16
    local.get 0
    local.get 2
    i32.store offset=12
    local.get 0
    local.get 1
    i32.store offset=8
    local.get 0
    i64.const -1
    i64.store)
  (func (;27;) (type 4) (param i32 i32)
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
    i32.const 65280
    i32.and
    i32.const 8
    i32.shl
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
    call 4
    drop
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;28;) (type 3) (param i32 i32) (result i32)
    (local i32)
    call 22
    local.tee 2
    local.get 0
    local.get 1
    call 14
    drop
    local.get 2)
  (func (;29;) (type 0)
    call 6
    i32.eqz
    if  ;; label = @1
      return
    end
    i32.const 131116
    i32.const 37
    call 7
    unreachable)
  (func (;30;) (type 5) (param i32 i32 i32) (result i32)
    local.get 0
    call 22
    local.tee 0
    call 8
    drop
    local.get 0
    call 9
    i32.const 32
    i32.ne
    if  ;; label = @1
      i32.const 131090
      i32.const 23
      call 28
      local.tee 0
      local.get 1
      local.get 2
      call 4
      drop
      local.get 0
      i32.const 131113
      i32.const 3
      call 4
      drop
      local.get 0
      i32.const 131326
      i32.const 16
      call 4
      drop
      local.get 0
      call 5
      unreachable
    end
    local.get 0)
  (func (;31;) (type 1) (param i32)
    call 10
    local.get 0
    i32.eq
    if  ;; label = @1
      return
    end
    i32.const 131153
    i32.const 25
    call 7
    unreachable)
  (func (;32;) (type 7) (param i32 i32 i32 i32)
    local.get 0
    local.get 1
    local.get 2
    call 33
    local.get 0
    i32.const 131408
    i32.const 131412
    local.get 3
    select
    i32.const 4
    i32.const 5
    local.get 3
    select
    call 33)
  (func (;33;) (type 8) (param i32 i32 i32)
    local.get 0
    local.get 1
    local.get 2
    call 28
    call 27)
  (func (;34;) (type 2) (result i32)
    (local i32)
    call 22
    local.tee 0
    i64.const 0
    call 0
    local.get 0)
  (func (;35;) (type 4) (param i32 i32)
    call 25
    drop
    local.get 0
    local.get 1
    call 23
    call 27)
  (func (;36;) (type 1) (param i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 2
    global.set 0
    local.get 0
    i32.load offset=16
    local.tee 4
    if  ;; label = @1
      call 22
      local.tee 1
      call 11
      i32.const 131316
      i32.const 10
      call 28
      local.tee 7
      local.get 1
      call 2
      drop
      local.get 2
      block (result i32)  ;; label = @2
        i32.const 141440
        i32.load8_u
        local.tee 1
        i32.eqz
        if  ;; label = @3
          i32.const 141440
          i32.const 1
          i32.store8
          i32.const 141436
          i32.const 0
          i32.store
          call 25
          br 1 (;@2;)
        end
        i32.const 131408
        i32.const 0
        call 28
      end
      i32.store offset=16
      local.get 2
      local.get 1
      i32.const 1
      i32.xor
      i32.store8 offset=20
      local.get 0
      i32.const 20
      i32.add
      i32.load
      local.tee 3
      local.get 2
      i32.const 16
      i32.add
      local.tee 1
      call 37
      local.get 1
      local.get 4
      local.get 3
      call 38
      local.get 0
      i32.const 24
      i32.add
      i32.load
      local.tee 6
      call 39
      local.get 1
      call 37
      local.get 6
      call 9
      local.set 8
      i32.const 0
      local.set 1
      loop  ;; label = @2
        local.get 8
        local.get 1
        i32.const 4
        i32.add
        local.tee 4
        i32.ge_u
        if  ;; label = @3
          local.get 2
          i32.const 0
          i32.store offset=24
          local.get 6
          local.get 1
          i32.const 4
          local.get 2
          i32.const 24
          i32.add
          local.tee 5
          call 12
          drop
          local.get 2
          local.get 2
          i32.load offset=24
          local.tee 1
          i32.const 24
          i32.shl
          local.get 1
          i32.const 65280
          i32.and
          i32.const 8
          i32.shl
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
          local.tee 1
          call 9
          local.tee 3
          i32.const 24
          i32.shl
          local.get 3
          i32.const 65280
          i32.and
          i32.const 8
          i32.shl
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
          i32.store offset=24
          local.get 2
          i32.const 16
          i32.add
          local.get 5
          i32.const 4
          call 38
          block  ;; label = @4
            local.get 2
            i32.load8_u offset=20
            if  ;; label = @5
              local.get 1
              call 9
              local.tee 5
              i32.const 10000
              i32.const 141436
              i32.load
              local.tee 3
              i32.sub
              i32.gt_u
              br_if 1 (;@4;)
              local.get 2
              i32.const 8
              i32.add
              local.get 3
              local.get 3
              local.get 5
              i32.add
              local.tee 3
              call 40
              local.get 1
              i32.const 0
              local.get 2
              i32.load offset=12
              local.get 2
              i32.load offset=8
              call 12
              drop
              i32.const 141436
              local.get 3
              i32.store
              local.get 4
              local.set 1
              br 3 (;@2;)
            end
            local.get 2
            i32.load offset=16
            local.get 1
            call 2
            drop
            local.get 4
            local.set 1
            br 2 (;@2;)
          end
          local.get 2
          i32.const 16
          i32.add
          call 41
          local.get 2
          i32.load offset=16
          local.get 1
          call 2
          drop
          local.get 4
          local.set 1
          br 1 (;@2;)
        end
      end
      local.get 2
      i32.load offset=16
      local.set 1
      local.get 2
      local.get 2
      i32.load8_u offset=20
      i32.store8 offset=28
      local.get 2
      local.get 1
      i32.store offset=24
      local.get 2
      i32.const 24
      i32.add
      call 41
      local.get 2
      i32.load offset=24
      local.set 1
      local.get 2
      i32.load8_u offset=28
      if  ;; label = @2
        i32.const 141436
        i32.const 0
        i32.store
        i32.const 141440
        i32.const 0
        i32.store8
      end
      local.get 7
      local.get 1
      call 13
      drop
    end
    local.get 0
    i32.load
    local.get 0
    i32.load offset=4
    local.get 0
    i32.load offset=8
    local.get 0
    i32.load offset=12
    call 15
    unreachable)
  (func (;37;) (type 4) (param i32 i32)
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
    i32.const 65280
    i32.and
    i32.const 8
    i32.shl
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
    call 38
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;38;) (type 8) (param i32 i32 i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 6
    global.set 0
    block  ;; label = @1
      block  ;; label = @2
        local.get 0
        i32.load8_u offset=4
        if  ;; label = @3
          i32.const 10000
          i32.const 141436
          i32.load
          local.tee 3
          i32.sub
          local.get 2
          i32.lt_u
          br_if 1 (;@2;)
          local.get 6
          i32.const 8
          i32.add
          local.get 3
          local.get 2
          local.get 3
          i32.add
          local.tee 8
          call 40
          local.get 6
          i32.load offset=8
          local.set 0
          block  ;; label = @4
            local.get 2
            local.get 6
            i32.load offset=12
            local.tee 3
            i32.eq
            if  ;; label = @5
              local.get 3
              i32.const 15
              i32.gt_u
              if  ;; label = @6
                local.get 0
                i32.const 0
                local.get 0
                i32.sub
                i32.const 3
                i32.and
                local.tee 5
                i32.add
                local.set 4
                local.get 5
                if  ;; label = @7
                  local.get 1
                  local.set 2
                  loop  ;; label = @8
                    local.get 0
                    local.get 2
                    i32.load8_u
                    i32.store8
                    local.get 2
                    i32.const 1
                    i32.add
                    local.set 2
                    local.get 0
                    i32.const 1
                    i32.add
                    local.tee 0
                    local.get 4
                    i32.lt_u
                    br_if 0 (;@8;)
                  end
                end
                local.get 4
                local.get 3
                local.get 5
                i32.sub
                local.tee 3
                i32.const -4
                i32.and
                local.tee 7
                i32.add
                local.set 0
                block  ;; label = @7
                  local.get 1
                  local.get 5
                  i32.add
                  local.tee 5
                  i32.const 3
                  i32.and
                  local.tee 2
                  if  ;; label = @8
                    local.get 7
                    i32.const 0
                    i32.le_s
                    br_if 1 (;@7;)
                    local.get 5
                    i32.const -4
                    i32.and
                    local.tee 9
                    i32.const 4
                    i32.add
                    local.set 1
                    i32.const 0
                    local.get 2
                    i32.const 3
                    i32.shl
                    local.tee 10
                    i32.sub
                    i32.const 24
                    i32.and
                    local.set 11
                    local.get 9
                    i32.load
                    local.set 2
                    loop  ;; label = @9
                      local.get 4
                      local.get 2
                      local.get 10
                      i32.shr_u
                      local.get 1
                      i32.load
                      local.tee 2
                      local.get 11
                      i32.shl
                      i32.or
                      i32.store
                      local.get 1
                      i32.const 4
                      i32.add
                      local.set 1
                      local.get 4
                      i32.const 4
                      i32.add
                      local.tee 4
                      local.get 0
                      i32.lt_u
                      br_if 0 (;@9;)
                    end
                    br 1 (;@7;)
                  end
                  local.get 7
                  i32.const 0
                  i32.le_s
                  br_if 0 (;@7;)
                  local.get 5
                  local.set 1
                  loop  ;; label = @8
                    local.get 4
                    local.get 1
                    i32.load
                    i32.store
                    local.get 1
                    i32.const 4
                    i32.add
                    local.set 1
                    local.get 4
                    i32.const 4
                    i32.add
                    local.tee 4
                    local.get 0
                    i32.lt_u
                    br_if 0 (;@8;)
                  end
                end
                local.get 3
                i32.const 3
                i32.and
                local.set 3
                local.get 5
                local.get 7
                i32.add
                local.set 1
              end
              local.get 3
              if  ;; label = @6
                local.get 0
                local.get 3
                i32.add
                local.set 2
                loop  ;; label = @7
                  local.get 0
                  local.get 1
                  i32.load8_u
                  i32.store8
                  local.get 1
                  i32.const 1
                  i32.add
                  local.set 1
                  local.get 0
                  i32.const 1
                  i32.add
                  local.tee 0
                  local.get 2
                  i32.lt_u
                  br_if 0 (;@7;)
                end
              end
              br 1 (;@4;)
            end
            call 42
            unreachable
          end
          i32.const 141436
          local.get 8
          i32.store
          br 2 (;@1;)
        end
        local.get 0
        i32.load
        local.get 1
        local.get 2
        call 4
        drop
        br 1 (;@1;)
      end
      local.get 0
      call 41
      local.get 0
      i32.load
      local.get 1
      local.get 2
      call 4
      drop
    end
    local.get 6
    i32.const 16
    i32.add
    global.set 0)
  (func (;39;) (type 6) (param i32) (result i32)
    local.get 0
    call 9
    i32.const 2
    i32.shr_u)
  (func (;40;) (type 8) (param i32 i32 i32)
    block  ;; label = @1
      local.get 1
      local.get 2
      i32.le_u
      if  ;; label = @2
        local.get 2
        i32.const 10000
        i32.gt_u
        br_if 1 (;@1;)
        local.get 0
        local.get 2
        local.get 1
        i32.sub
        i32.store offset=4
        local.get 0
        local.get 1
        i32.const 131436
        i32.add
        i32.store
        return
      end
      call 42
      unreachable
    end
    call 42
    unreachable)
  (func (;41;) (type 1) (param i32)
    (local i32)
    local.get 0
    i32.load8_u offset=4
    local.set 1
    local.get 0
    i32.const 0
    i32.store8 offset=4
    block  ;; label = @1
      local.get 1
      if  ;; label = @2
        i32.const 141436
        i32.load
        local.tee 1
        i32.const 10001
        i32.ge_u
        br_if 1 (;@1;)
        local.get 0
        i32.load
        i32.const 131436
        local.get 1
        call 4
        drop
        i32.const 141436
        i32.const 0
        i32.store
        i32.const 141440
        i32.const 0
        i32.store8
      end
      return
    end
    call 42
    unreachable)
  (func (;42;) (type 0)
    call 50
    unreachable)
  (func (;43;) (type 1) (param i32)
    (local i32 i32)
    i32.const -11
    local.set 1
    block  ;; label = @1
      i32.const 141444
      i32.load8_u
      local.tee 2
      if  ;; label = @2
        i32.const -11
        i32.const 2147483647
        local.get 2
        select
        local.set 1
        br 1 (;@1;)
      end
      i32.const 141444
      i32.const 1
      i32.store8
      i32.const -11
      call 16
    end
    local.get 0
    local.get 1
    call 24)
  (func (;44;) (type 0)
    call 17
    i32.const 0
    call 31)
  (func (;45;) (type 0)
    call 29
    i32.const 1
    call 31
    i32.const 0
    i32.const 131376
    i32.const 7
    call 30
    call 43)
  (func (;46;) (type 0)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32)
    global.get 0
    i32.const 80
    i32.sub
    local.tee 1
    global.set 0
    call 29
    i32.const 1
    call 31
    i32.const 0
    i32.const 131376
    i32.const 7
    call 30
    call 22
    local.tee 0
    i64.const 1000
    call 0
    local.get 0
    call 24
    call 22
    local.tee 3
    i64.const 50000000000000000
    call 0
    i32.const 131342
    i32.const 5
    call 28
    local.set 4
    i32.const 131342
    i32.const 5
    call 28
    local.set 5
    local.get 1
    i32.const 1
    i32.store8 offset=78
    local.get 1
    i32.const 0
    i32.store16 offset=76
    local.get 1
    i32.const 16843009
    i32.store offset=72
    global.get 0
    i32.const 32
    i32.sub
    local.tee 2
    global.set 0
    call 34
    drop
    local.get 1
    i32.const 72
    i32.add
    local.tee 0
    i32.load8_u offset=6
    local.set 6
    local.get 0
    i32.load8_u offset=5
    local.set 7
    local.get 0
    i32.load8_u offset=4
    local.set 8
    local.get 0
    i32.load8_u offset=3
    local.set 9
    local.get 0
    i32.load8_u offset=2
    local.set 10
    local.get 0
    i32.load8_u offset=1
    local.set 11
    local.get 0
    i32.load8_u
    local.set 12
    local.get 2
    i32.const 131268
    i32.const 32
    call 28
    i32.const 131300
    i32.const 16
    call 26
    local.get 2
    i32.const 24
    i32.add
    local.tee 13
    local.get 3
    i32.store
    local.get 2
    i32.const 16
    i32.add
    local.tee 3
    i32.load
    local.tee 0
    local.get 4
    call 35
    local.get 0
    local.get 5
    call 35
    local.get 0
    i32.const 131178
    i32.const 9
    local.get 12
    call 32
    local.get 0
    i32.const 131187
    i32.const 7
    local.get 11
    call 32
    local.get 0
    i32.const 131194
    i32.const 8
    local.get 10
    call 32
    local.get 0
    i32.const 131202
    i32.const 24
    local.get 9
    call 32
    local.get 0
    i32.const 131226
    i32.const 14
    local.get 8
    call 32
    local.get 0
    i32.const 131240
    i32.const 10
    local.get 7
    call 32
    local.get 0
    i32.const 131250
    i32.const 18
    local.get 6
    call 32
    local.get 1
    i32.const 40
    i32.add
    local.tee 0
    i32.const 24
    i32.add
    local.get 13
    i64.load
    i64.store
    local.get 0
    i32.const 16
    i32.add
    local.get 3
    i64.load
    i64.store
    local.get 0
    i32.const 8
    i32.add
    local.get 2
    i32.const 8
    i32.add
    i64.load
    i64.store
    local.get 0
    local.get 2
    i64.load
    i64.store
    local.get 2
    i32.const 32
    i32.add
    global.set 0
    local.get 1
    i32.const 0
    i32.store offset=24
    local.get 1
    local.get 1
    i64.load offset=52 align=4
    i64.store offset=16
    local.get 1
    local.get 1
    i32.load offset=64
    i32.store offset=12
    local.get 1
    local.get 1
    i32.load offset=48
    i32.store offset=8
    local.get 1
    i32.const 8
    i32.add
    call 36
    unreachable)
  (func (;47;) (type 0)
    (local i32 i32 i32 i32 i32 i64)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 0
    global.set 0
    call 29
    i32.const 2
    call 31
    i32.const 0
    i32.const 131361
    i32.const 15
    call 30
    local.set 1
    i32.const 1
    i32.const 131347
    i32.const 14
    call 30
    call 34
    local.set 3
    call 25
    local.set 4
    call 18
    local.get 0
    i32.const 5
    i32.store16 offset=14
    local.get 3
    local.get 1
    local.get 0
    i32.const 14
    i32.add
    i32.const 2
    call 19
    local.get 4
    call 22
    call 22
    local.tee 1
    call 20
    drop
    call 21
    local.get 1
    call 39
    drop
    call 43
    local.get 0
    i32.const 16
    i32.add
    global.set 0)
  (func (;48;) (type 0)
    (local i32 i32 i32 i32)
    global.get 0
    i32.const -64
    i32.add
    local.tee 1
    global.set 0
    call 17
    i32.const 2
    call 31
    i32.const 0
    i32.const 131392
    i32.const 16
    call 30
    i32.const 1
    i32.const 131383
    i32.const 9
    call 30
    local.set 2
    call 23
    local.set 3
    global.get 0
    i32.const 32
    i32.sub
    local.tee 0
    global.set 0
    local.get 0
    i32.const 8
    i32.add
    local.get 3
    i32.const 131072
    i32.const 18
    call 26
    call 25
    drop
    local.get 2
    call 23
    local.set 2
    local.get 0
    i32.const 24
    i32.add
    local.tee 3
    i32.load
    local.get 2
    call 27
    local.get 1
    i32.const 40
    i32.add
    local.tee 2
    i32.const 16
    i32.add
    local.get 3
    i64.load
    i64.store
    local.get 2
    i32.const 8
    i32.add
    local.get 0
    i32.const 16
    i32.add
    i64.load
    i64.store
    local.get 2
    local.get 0
    i64.load offset=8
    i64.store
    local.get 0
    i32.const 32
    i32.add
    global.set 0
    call 34
    local.set 0
    local.get 1
    i32.const 0
    i32.store offset=24
    local.get 1
    local.get 0
    i32.store offset=12
    local.get 1
    local.get 1
    i64.load offset=52 align=4
    i64.store offset=16
    local.get 1
    local.get 1
    i32.load offset=48
    i32.store offset=8
    local.get 1
    i32.const 8
    i32.add
    call 36
    unreachable)
  (func (;49;) (type 0)
    nop)
  (func (;50;) (type 0)
    i32.const 131417
    i32.const 14
    call 7
    unreachable)
  (memory (;0;) 3)
  (global (;0;) (mut i32) (i32.const 131072))
  (global (;1;) i32 (i32.const 141445))
  (global (;2;) i32 (i32.const 141456))
  (export "memory" (memory 0))
  (export "init" (func 44))
  (export "send_rewa" (func 45))
  (export "send_and_issue_token" (func 46))
  (export "do_deploy" (func 47))
  (export "change_owner" (func 48))
  (export "callBack" (func 49))
  (export "__data_end" (global 1))
  (export "__heap_base" (global 2))
  (data (;0;) (i32.const 131072) "ChangeOwnerAddressargument decode error (): function does not accept DCDT paymentwrong number of argumentscanFreezecanWipecanPausecanTransferNFTCreateRolecanChangeOwnercanUpgradecanAddSpecialRoles")
  (data (;1;) (i32.const 131277) "\01")
  (data (;2;) (i32.const 131297) "\02\ff\ffissueNonFungibleCB_CLOSUREbad array lengthMACIUsecond_addresssource_contractaddressnew_ownerchild_sc_addresstruefalsepanic occurred")
  (data (;3;) (i32.const 131432) "\9c\ff\ff\ff"))
