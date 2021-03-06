import $axios from '@/utils/axios'

export function getStats (name: string, address: string) {
    return $axios.get(`/api/v1/b/stats?service=${name}&address=${address}`)
}

export function getAPIStats (name: string, address: string) {
    return $axios.get(`/api/v1/b/api-stats?name=${name}&address=${address}`)
}
