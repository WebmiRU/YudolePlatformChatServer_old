class APIService {
    async getModules() {
        const response = await fetch('http://127.0.0.1/api/modules')
        return await response.json();
    }

    async getModulesId(id: string) {
        const response = await fetch('http://127.0.0.1/api/modules/' + id)
        return await response.json();
    }
}

export default new APIService