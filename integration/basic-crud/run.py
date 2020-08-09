import unittest
import urllib
import urllib.request
import urllib.error
import subprocess
import json

class BasicCRUDTest(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        subprocess.run('bash ./integration/utils/start-server.sh ./integration/basic-crud/api.yml'.split(' '))

    def test_get(self):
        response = urllib.request.urlopen('http://localhost:8000/students/1').read().decode('utf-8')
        student = json.loads(response)
        self.assertEqual(student['name'], 'A')
        self.assertEqual(student['school_id'], 1)
        self.assertEqual(student['birthday'], '1990-01-01T00:00:00Z')

    def test_scan(self):
        response = urllib.request.urlopen('http://localhost:8000/students?limit=100&offset=0').read().decode('utf-8')
        students = json.loads(response)
        self.assertEqual(len(students), 4)

    def test_post(self):
        response = urllib.request.urlopen('http://localhost:8000/students?limit=100&offset=0').read().decode('utf-8')
        old_count = len(json.loads(response))

        data = json.dumps({'name': 'E', 'birthday': '1990-05-05T00:00:00Z', 'school_id': 5})
        request = urllib.request.Request(url='http://localhost:8000/students', data=data.encode(encoding='utf-8'), method='POST')
        response = urllib.request.urlopen(request).read().decode('utf-8')
        student_post_result = json.loads(response)
        print('+++++\n')
        print(student_post_result)
        print(student_post_result['id'])
        print('+++++\n')
        response = urllib.request.urlopen('http://localhost:8000/students/{}'.format(student_post_result['id'])).read().decode('utf-8')
        student_get_result = json.loads(response)
        self.assertDictEqual(student_post_result, student_get_result)

        response = urllib.request.urlopen('http://localhost:8000/students?limit=100&offset=0').read().decode('utf-8')
        new_count = len(json.loads(response))
        self.assertEqual(new_count, old_count+1)

    def test_put(self):
        response = urllib.request.urlopen('http://localhost:8000/students?limit=100&offset=0').read().decode('utf-8')
        old_count = len(json.loads(response))

        new_info = {'id': 1, 'name': 'F', 'birthday': '1990-06-06T00:00:00Z', 'school_id': 6}
        data = json.dumps(new_info)
        request = urllib.request.Request(url='http://localhost:8000/students/1', data=data.encode(encoding='utf-8'), method='PUT')
        urllib.request.urlopen(request)
        response = urllib.request.urlopen('http://localhost:8000/students/1').read().decode('utf-8')
        student = json.loads(response)
        self.assertEqual(student, new_info)

        response = urllib.request.urlopen('http://localhost:8000/students?limit=100&offset=0').read().decode('utf-8')
        new_count = len(json.loads(response))
        self.assertEqual(new_count, old_count)

    def test_delete(self):
        response = urllib.request.urlopen('http://localhost:8000/students?limit=100&offset=0').read().decode('utf-8')
        old_count = len(json.loads(response))

        request = urllib.request.Request(url='http://localhost:8000/students/2', method='DELETE')
        urllib.request.urlopen(request)
        with self.assertRaises(urllib.error.HTTPError) as err:
            urllib.request.urlopen('http://localhost:8000/students/2').read().decode('utf-8')
        self.assertEqual(err.exception.code, 404)

        response = urllib.request.urlopen('http://localhost:8000/students?limit=100&offset=0').read().decode('utf-8')
        new_count = len(json.loads(response))
        self.assertEqual(new_count, old_count-1)


if __name__ == '__main__':
    loader = unittest.TestLoader()
    loader.sortTestMethodsUsing = None
    unittest.main(testLoader=loader)