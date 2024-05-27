import {reactive, ref} from "vue";

export default reactive({
    breadcrumbs: [
        { label: 'Components' },
        { label: 'Form' },
        { label: 'InputText', route: '/inputtext' }],
})