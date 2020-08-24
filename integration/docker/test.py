import os
import subprocess
import unittest


# todo: check whether the dockerfile generated is right, ie. can build and work
class DockerfileGeneratingTest(unittest.TestCase):
    def test_default(self):
        subprocess.run("./autoAPI --force -f ./integration/docker/default.yml --force -o ./api".split(' '))
        self.assertTrue(os.path.exists("./api/Dockerfile"))

    def test_command_line(self):
        # force noDocker
        subprocess.run(
            "./autoAPI --force -f ./integration/docker/default.yml --force --nodocker -o ./api".split(' '))
        self.assertTrue(os.path.exists("./api"))
        self.assertFalse(os.path.exists("./api/Dockerfile"))


if __name__ == '__main__':
    loader = unittest.TestLoader()
    loader.sortTestMethodsUsing = None
    unittest.main(testLoader=loader)
