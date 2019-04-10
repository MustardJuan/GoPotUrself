package shell

func Df() string {
	output := ("Filesystem     1K-blocks    Used Available Use% Mounted on\n" +
    "overlay         61255492 1886312  56227856   4% /\n" +
    "tmpfs              65536       0     65536   0% /dev\n" +
    "tmpfs            1023372       0   1023372   0% /sys/fs/cgroup\n"+
    "/dev/sda1       61255492 1886312  56227856   4% /etc/hosts\n" +
    "shm                65536       0     65536   0% /dev/shm\n" +
    "tmpfs            1023372       0   1023372   0% /proc/acpi\n" +
    "tmpfs            1023372       0   1023372   0% /sys/firmware\n")
	
    return output
}