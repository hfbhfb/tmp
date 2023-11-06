[root@pprmcce01 ~]# fio --randrepeat=1 --ioengine=libaio --direct=1 --gtod_reduce=1 --name=testa --filename=randon_read_write.fio --bs=4k --iodepth=64 --size=1G --readwrite=randrw --rwmixread=75
testa: (g=0): rw=randrw, bs=(R) 4096B-4096B, (W) 4096B-4096B, (T) 4096B-4096B, ioengine=libaio, iodepth=64
fio-3.7
Starting 1 process
testa: Laying out IO file (1 file / 1024MiB)
Jobs: 1 (f=1): [m(1)][100.0%][r=73.0MiB/s,w=24.7MiB/s][r=18.7k,w=6316 IOPS][eta 00m:00s]
testa: (groupid=0, jobs=1): err= 0: pid=208336: Mon Nov  6 14:30:07 2023
   read: IOPS=8151, BW=31.8MiB/s (33.4MB/s)(768MiB/24105msec)
   bw (  KiB/s): min= 1248, max=100032, per=99.15%, avg=32328.92, stdev=16654.30, samples=48
   iops        : min=  312, max=25008, avg=8082.21, stdev=4163.56, samples=48
  write: IOPS=2723, BW=10.6MiB/s (11.2MB/s)(256MiB/24105msec)
   bw (  KiB/s): min=  496, max=33640, per=99.15%, avg=10800.46, stdev=5609.33, samples=48
   iops        : min=  124, max= 8410, avg=2700.08, stdev=1402.31, samples=48
  cpu          : usr=0.78%, sys=3.51%, ctx=35376, majf=0, minf=6
  IO depths    : 1=0.1%, 2=0.1%, 4=0.1%, 8=0.1%, 16=0.1%, 32=0.1%, >=64=100.0%
     submit    : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     complete  : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.1%, >=64=0.0%
     issued rwts: total=196498,65646,0,0 short=0,0,0,0 dropped=0,0,0,0
     latency   : target=0, window=0, percentile=100.00%, depth=64

Run status group 0 (all jobs):
   READ: bw=31.8MiB/s (33.4MB/s), 31.8MiB/s-31.8MiB/s (33.4MB/s-33.4MB/s), io=768MiB (805MB), run=24105-24105msec
  WRITE: bw=10.6MiB/s (11.2MB/s), 10.6MiB/s-10.6MiB/s (11.2MB/s-11.2MB/s), io=256MiB (269MB), run=24105-24105msec

Disk stats (read/write):
    dm-0: ios=192121/64084, merge=0/0, ticks=1291988/254708, in_queue=1546696, util=100.00%, aggrios=196838/65660, aggrmerge=0/0, aggrticks=1300922/257175, aggrin_queue=1408008, aggrutil=100.00%
  sda: ios=196838/65660, merge=0/0, ticks=1300922/257175, in_queue=1408008, util=100.00%

[root@zdmdb01 ~]# fio --randrepeat=1 --ioengine=libaio --direct=1 --gtod_reduce=1 --name=testa --filename=randon_read_write.fio --bs=4k --iodepth=64 --size=1G --readwrite=randrw --rwmixread=75
testa: (g=0): rw=randrw, bs=(R) 4096B-4096B, (W) 4096B-4096B, (T) 4096B-4096B, ioengine=libaio, iodepth=64
fio-3.7
Starting 1 process
testa: Laying out IO file (1 file / 1024MiB)
Jobs: 1 (f=1): [m(1)][93.8%][r=71.1MiB/s,w=24.1MiB/s][r=18.2k,w=6170 IOPS][eta 00m:01s]
testa: (groupid=0, jobs=1): err= 0: pid=2259593: Mon Nov  6 16:39:20 2023
   read: IOPS=12.5k, BW=48.8MiB/s (51.2MB/s)(768MiB/15722msec)
   bw (  KiB/s): min=27616, max=101496, per=99.79%, avg=49887.00, stdev=16857.71, samples=31
   iops        : min= 6904, max=25374, avg=12471.74, stdev=4214.45, samples=31
  write: IOPS=4175, BW=16.3MiB/s (17.1MB/s)(256MiB/15722msec)
   bw (  KiB/s): min= 9016, max=34280, per=99.81%, avg=16669.55, stdev=5745.89, samples=31
   iops        : min= 2254, max= 8570, avg=4167.39, stdev=1436.47, samples=31
  cpu          : usr=3.03%, sys=13.42%, ctx=140395, majf=0, minf=106
  IO depths    : 1=0.1%, 2=0.1%, 4=0.1%, 8=0.1%, 16=0.1%, 32=0.1%, >=64=100.0%
     submit    : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.0%, >=64=0.0%
     complete  : 0=0.0%, 4=100.0%, 8=0.0%, 16=0.0%, 32=0.0%, 64=0.1%, >=64=0.0%
     issued rwts: total=196498,65646,0,0 short=0,0,0,0 dropped=0,0,0,0
     latency   : target=0, window=0, percentile=100.00%, depth=64

Run status group 0 (all jobs):
   READ: bw=48.8MiB/s (51.2MB/s), 48.8MiB/s-48.8MiB/s (51.2MB/s-51.2MB/s), io=768MiB (805MB), run=15722-15722msec
  WRITE: bw=16.3MiB/s (17.1MB/s), 16.3MiB/s-16.3MiB/s (17.1MB/s-17.1MB/s), io=256MiB (269MB), run=15722-15722msec

Disk stats (read/write):
    dm-0: ios=196263/65561, merge=0/0, ticks=762228/213688, in_queue=975916, util=99.39%, aggrios=196498/65646, aggrmerge=0/0, aggrticks=779420/216049, aggrin_queue=941284, aggrutil=98.78%
  sda: ios=196498/65646, merge=0/0, ticks=779420/216049, in_queue=941284, util=98.78%