import os
import subprocess
import unittest


class K8sGeneratingTest(unittest.TestCase):
    def test_full(self):
        subprocess.run("./autoAPI --force -f ./integration/k8s/full.yml --force -o ./api".split(' '))
        self.assert_(os.path.exists("./api/student.yaml"))
        with open("./api/student.yaml", "r") as f:
            content = f.read()
            self.assertEqual(content.count("development"), 3)
            self.assertIn("/api/student-svc", content)
            self.assertIn("cloud.shuosc.com", content)

    def test_host_only(self):
        subprocess.run("./autoAPI --force -f ./integration/k8s/host-only.yml --force -o ./api1".split(' '))
        self.assert_(os.path.exists("./api1/student.yaml"))
        with open("./api1/student.yaml", "r") as f:
            content = f.read()
            self.assertNotIn("namespace", content)
            self.assertIn("/api/student", content)
            self.assertIn("cloud.shuosc.com", content)

    def test_no_ingress(self):
        subprocess.run("./autoAPI --force -f ./integration/k8s/no-ingress.yml --force -o ./api2".split(' '))
        self.assertTrue(os.path.exists("./api2/student.yaml"))
        with open("./api2/student.yaml", "r") as f:
            content = f.read()
            self.assertNotIn("Ingress", content)

    def test_not_generate(self):
        subprocess.run("./autoAPI --nok8s --force -f ./integration/k8s/no-ingress.yml --force -o ./api3".split(' '))
        self.assertFalse(os.path.exists("./api3/student.yaml"))


if __name__ == '__main__':
    unittest.main()
