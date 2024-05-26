class APIService {
    async getExtensions() {
        const response = await fetch('https://dummyjson.com/users')
        return await response.json();
    }
}

export default new APIService