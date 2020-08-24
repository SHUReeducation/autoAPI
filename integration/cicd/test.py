import os
import subprocess
import unittest


class CICDGeneratingTest(unittest.TestCase):
    def test_default(self):
        subprocess.run("./autoAPI --force -f ./integration/docker/default.yml --force -o ./api".split(' '))
        self.assertFalse(os.path.exists("./api/.github/workflows/dockerimage.yml"))

    def test_command_line(self):
        # force noDocker
        subprocess.run(
            "./autoAPI --force -f ./integration/docker/default.yml --force --nodocker -o ./api".split(' '))
        self.assertTrue(os.path.exists("./api"))
        self.assertFalse(os.path.exists("./api/.github/workflows/dockerimage.yml"))

        subprocess.run(
            "./autoAPI --force -f ./integration/docker/default.yml --force -nd -o ./api".split(' '))
        self.assertTrue(os.path.exists("./api"))
        self.assertFalse(os.path.exists("./api/.github/workflows/dockerimage.yml"))

    def test_docker_image_name(self):
        # set docker info by command
        subprocess.run(
            "./autoAPI --force -f ./integration/docker/default.yml --force -du testuser -dt v0 -o ./api".split(
                ' '))
        self.assertTrue(os.path.exists("./api/.github/workflows/dockerimage.yml"))
        with open("./api/.github/workflows/dockerimage.yml") as f:
            content = f.read()
            self.assertEqual(content.count("testuser/student:v0"), 2)
            self.assertIn('--username "testuser"', content)
        # set docker info by env
        subprocess.run(
            "./autoAPI --force -f ./integration/docker/default.yml --force -o ./api".split(' '),
            env={'DOCKER_USERNAME': 'testuser', 'DOCKER_TAG': 'v0'})
        with open("./api/.github/workflows/dockerimage.yml") as f:
            content = f.read()
            self.assertEqual(content.count("testuser/student:v0"), 2)
            self.assertIn('--username "testuser"', content)
        # simulate GitHub Actions env
        subprocess.run(
            "./autoAPI --force -f ./integration/docker/default.yml --force -o ./api".split(' '),
            env={'DOCKER_USERNAME': 'testuser', 'GITHUB_RUN_NUMBER': '7'})
        with open("./api/.github/workflows/dockerimage.yml") as f:
            content = f.read()
            self.assertEqual(content.count("testuser/student:ci-v7"), 2)
            self.assertIn('--username "testuser"', content)


if __name__ == '__main__':
    loader = unittest.TestLoader()
    loader.sortTestMethodsUsing = None
    unittest.main(testLoader=loader)
