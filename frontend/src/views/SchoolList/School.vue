<template>
    <v-app>
        <!-- <v-app-bar app class="app-bar" dark color="#ffffff">
            <v-toolbar-side-icon>
                <router-link to="/test">
                    <v-img height="60" width="" src="@/assets/cube.jpg" />
                </router-link>
            </v-toolbar-side-icon>

            <v-toolbar-title>
                <p class="mytitle">Trường học có tổng số sinh viên là : {{ totalStudent }}</p>
            </v-toolbar-title>
            <router-link to="/">
                <p class="mytitle-right">Bạn cần giúp đỡ?</p>
            </router-link>
            <v-avatar>
                <img height="30" src="@/assets/logo.png" alt="John" />
            </v-avatar>
        </v-app-bar> -->

        <v-main>
            <v-container class="container">
                <template v-if="errorAddStudentForm">
                    <v-alert time=5 icon="$vuetify" text="Lỗi không thể thêm sinh viên" type="error"
                        variant="tonal"></v-alert>
                </template>
                <template v-if="successAddStudentForm">
                    <v-alert icon="$vuetify" text="Đã thêm sinh viên thành công" type="success" variant="tonal"></v-alert>
                </template>

                <v-select v-model="selectedType" :items="types" label="Select Data Type" class="select"></v-select>

                <template v-if="selectedType === 'teachers'">
                    <h2>Teachers</h2>
                    <div class="table-container">
                        <table>
                            <caption>Bảng danh sách giáo viên</caption>
                            <thead>
                                <tr>
                                    <th>ID</th>
                                    <th>Name</th>
                                    <th>Subject</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr v-for="teacher in teachers" :key="teacher.id">
                                    <td>{{ teacher.ID }}</td>
                                    <td>{{ teacher.Name }}</td>
                                    <td>{{ teacher.Subject }}</td>
                                </tr>
                            </tbody>
                        </table>
                    </div>
                </template>

                <template v-if="selectedType === 'classes'">
                    <h2>Class</h2>
                    <div class="table-container">
                        <table>
                            <thead>
                                <tr>
                                    <th>ID</th>
                                    <th>Name</th>
                                    <th>TeacherID</th>
                                    <th>Student_Count</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr v-for="classs in classes" :key="classs.id">
                                    <td>{{ classs.ID }}</td>
                                    <td>{{ classs.Name }}</td>
                                    <td>{{ classs.TeacherID }}</td>
                                    <td>{{ classs.StudentCount }}</td>
                                </tr>
                            </tbody>
                        </table>
                    </div>
                </template>

                <template v-if="selectedType === 'students'">
                    <h2>Students</h2>


                    <div class="table-container">
                        <table>
                            <thead>
                                <tr>
                                    <th>ID</th>
                                    <th>Name</th>
                                    <th>ClassID</th>
                                    <th>Age</th>
                                    <th>Address</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr v-for="student in students.slice((currentPage - 1) * 10, currentPage * 10)"
                                    :key="student.id">
                                    <td>{{ student.ID }}</td>
                                    <td>{{ student.Name }}</td>
                                    <td>{{ student.ClassID }}</td>
                                    <td>{{ student.Age }}</td>
                                    <td>{{ student.Address }}</td>
                                </tr>
                            </tbody>
                        </table>

                        <v-pagination v-model="currentPage" :length="Math.ceil(students.length / 10)" :total-visible="5"
                            class="mt-4"></v-pagination>
                    </div>
                </template>

                <template v-if="selectedType === 'add student'">
                    <v-form ref="studentForm" v-model="valid">
                        <v-text-field v-model="newStudent.id" label="ID"></v-text-field>
                        <v-text-field v-model="newStudent.name" label="Name"></v-text-field>
                        <v-text-field v-model="newStudent.classId" label="Class ID"></v-text-field>
                        <v-text-field v-model="newStudent.age" label="Age"></v-text-field>
                        <v-text-field v-model="newStudent.address" label="Address"></v-text-field>
                    </v-form>
                    <v-btn color="primary" @click="addStudent">Add Student</v-btn>

                </template>
                <template v-if="selectedType === 'delete student'">
                    <v-form ref="abc" v-model="valid">
                        <v-text-field v-model="oldStudent.id" label="ID"></v-text-field>
                    </v-form>
                    <v-btn color="primary" @click="deleteStudent(oldStudent.id)">Delete Student</v-btn>
                </template>


            </v-container>
        </v-main>
    </v-app>
</template>
  
<script lang="ts" setup>
import { ref, watch, onMounted } from 'vue';
import axios from 'axios';
import { computed, toRef } from 'vue';

const selectedType = ref('');
const types = ['teachers', 'classes', 'students', 'add student', 'delete student'];
onMounted(() => {
    const stored = localStorage.getItem('selectedType');
    if (stored) {
        selectedType.value = stored;
    }
})
const teachers = ref<any[]>([]);
const classes = ref<any[]>([]);
const students = ref<any[]>([]);
const currentPage = ref(1);
const totalStudent = computed(() => {
    return students.value.length
})

const nextID = computed(() => {
    return students.value.length + 1
})

let key = ref(toRef(totalStudent));
let key1 = ref(toRef(nextID));
let showForm = ref(false)
let errorAddStudentForm = ref(false)
let successAddStudentForm = ref(false)
let notify = ref(false)
const newStudent = ref({
    id: key1,
    name: '',
    classId: 1,
    age: 1,
    address: ''
});


const oldStudent = ref({
    id: key,
});

const valid = ref(false);

const loadData = () => {
    notifyHidden();
    if (selectedType.value === 'teachers') {
        axios
            .get('http://127.0.0.1:3030/users/teachers', {})
            .then((response: { data: any[] }) => {
                teachers.value = response.data;
            })
            .catch((error: any) => {
                console.error(error);
            });
        console.log("teacher")
    } else if (selectedType.value === 'classes') {
        axios
            .get('http://127.0.0.1:3030/users/classes', {})
            .then((response: { data: any[] }) => {
                classes.value = response.data;
            })
            .catch((error: any) => {
                console.error(error);
            });
    } else if (selectedType.value === 'students') {
        axios
            .get('http://127.0.0.1:3030/users/students', {})
            .then((response: { data: any[] }) => {
                students.value = response.data;
            })
            .catch((error: any) => {
                console.error(error);
            });
    }

};

const notifyHidden = () => {
    successAddStudentForm.value = false;
    errorAddStudentForm.value = false;
}

const addStudent = () => {
    axios
        .post('http://127.0.0.1:3030/users/students', newStudent.value)
        .then((response) => {
            console.log("Thanh cong ", newStudent.value)
            loadData();
            if (notify.value === false) successAddStudentForm.value = true;
            else {
                notifyHidden();
                successAddStudentForm.value = true;
                notify.value = true;
            }
            newStudent.value = {
                id: key1.value,
                name: '',
                classId: 1,
                age: 1,
                address: ''
            };
            showForm.value = false;
        })
        .catch((error) => {
            console.log("that bat ", newStudent.value)
            console.error(error);
            if (notify.value === false) errorAddStudentForm.value = true;
            else {
                notifyHidden();
                errorAddStudentForm.value = true;
                notify.value = true;
            }
        });
}

const deleteStudent = (studentID: any) => {
    axios
        .delete(`http://127.0.0.1:3030/users/students/${studentID}`)
        .then(response => () => {
            loadData();
            reloadPage();
            console.log("thanh cong", studentID)
        })
        .catch((error) => {
            console.log("that bat ", studentID)
            console.error(error);
        });
}

const reloadPage = () => {
    location.reload();
}


loadData();

watch(selectedType, (newValue) => {
    loadData();
    localStorage.setItem('selectedType', newValue)

})

    ;


</script>


<style lang="scss">
// Thiết lập các biến tùy chỉnh cho màu sắc và kích thước
$primary-color: #2196f3;
$secondary-color: #ff4081;
$background-color: #f5f5f5;
$font-color: #333;
$container-width: 1200px;

// Thiết lập các biến tùy chỉnh cho các thành phần khác nhau
$app-bar-height: 80px;
$heading-font-size: 24px;
$paragraph-font-size: 16px;

.app-bar {
    height: $app-bar-height;
    background-color: $primary-color;
    color: white;
}

.select {
    margin-top: 0px;
}

.container {
    max-width: $container-width;
    padding: 200px;

    .mt-4 {
        display: flex;
        flex-wrap: wrap;
        justify-content: center;
        align-items: center;
    }

    table {
        width: 100%;
        border-collapse: collapse;
        margin-bottom: 20px;
        font-size: 17px;
    }

    th,
    td {
        padding: 8px;
        text-align: left;
        border-bottom: 1px solid #ddd;
    }

    th {
        background-color: $primary-color;
        color: white;
    }


    @media (max-width: 620px) {
        .app-bar {

            .mytitle {
                font-size: 12px;
            }
        }

        .container {
            padding: 10px;
        }
    }

    .table {
        font-size: 12px;
    }
}

@media (max-width: 400px) {
    .app-bar {
        .mytitle {
            font-size: 8px;
        }
    }

    .container {
        padding: 5px;
    }

    table {
        font-size: 10px;
    }
}</style>