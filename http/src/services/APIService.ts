class APIService {
    async modulesIndexGet() {
        const response = await fetch('http://127.0.0.1/api/modules')
        return await response.json();
    }

    async modulesIdGet(id: string) {
        const response = await fetch('http://127.0.0.1/api/modules/' + id)
        return await response.json();
    }

    async modulesIdPut(id: string, payload: object) {
        const response = await fetch('http://127.0.0.1/api/modules/' + id, {
            headers: {
                'Content-type': 'application/json'
            },
            method: 'PUT',
            body: JSON.stringify(payload),
        })
        return await response.json();
    }

    async putModulesIdSetState(id: string, state: string) {
        const response = await fetch('http://127.0.0.1/api/modules/' + id + '/state/' + state, {
            method: 'PUT',
        })
        return await response.json();
    }

    async modulesIdStart(id: string) {
        const response = await fetch('http://127.0.0.1/api/modules/' + id + '/start', {
            method: 'POST',
        })
        return await response.json();
    }

    async modulesIdStop(id: string) {
        const response = await fetch('http://127.0.0.1/api/modules/' + id + '/stop', {
            method: 'POST',
        })
        return await response.json();
    }
}

export default new APIService