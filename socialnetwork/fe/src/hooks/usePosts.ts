import React, { useMemo } from "react"
import { useFetch } from "./useFetch"


export interface IPost {
    id: number
    createdAt: number
    authorId: number
    author: string
    threadId: string
    op: number,
    content: string
}

export const usePosts = () => {
    return useFetch<IPost[]>("http://localhost:8080/posts")
}

export const useThreads = () => {
    const { data, loading, error } = useFetch<IPost[]>("http://localhost:8080/posts")

    const filtered = useMemo(() => {
        return data ? data.filter(p => p.op === 1) : []
    }, [data])
    return {
        threads: filtered,
        loading,
        error
    }
}

export const useUserPosts = (userId: number) => {
    const { data, loading, error } = useFetch<IPost[]>("http://localhost:8080/posts")

    const filtered = useMemo(() => {
        return data ? data.filter(p => p.authorId === userId) : []
    }, [data, userId])

    return {
        posts: filtered,
        loading,
        error
    }
}

export const useThreadPosts = (threadId: string) => {
    const { data, loading, error } = useFetch<IPost[]>("http://localhost:8080/posts")

    const filtered = useMemo(() => {
        return data ? data.filter(p => p.threadId == threadId) : []
    }, [data, threadId])

    return {
        posts: filtered,
        loading,
        error
    }

}