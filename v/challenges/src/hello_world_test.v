module hello_world_test

import hello_world

fn test_hello_world() {
	assert hello_world.hello_world() == 'hello world'
}
