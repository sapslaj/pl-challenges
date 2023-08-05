const std = @import("std");
const expect = @import("std").testing.expect;

fn helloWorld() []const u8 {
    return "hello world";
}

test "hello world" {
    try expect(std.mem.eql(u8, "hello world", helloWorld()));
}
