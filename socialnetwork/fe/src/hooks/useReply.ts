import React from 'react'
import { IPost } from './usePosts'
import { IProfile } from './useProfiles'


interface IPostReplyRequest {
    authorId: number,
    threadId: string,
    content: string
}

interface ICreateThreadRequest {
    authorId: number,
    content: string
}

export const usePosting = () => {

    const replyToThread = (
        user: IProfile,
        message: string,
        post: IPost,
        callBack?: (result: string) => void,
        errorCallback?: (reason: string) => void
    ) => {
        const request: IPostReplyRequest = {
            authorId: user.id,
            threadId: post.threadId,
            content: message
        }

        fetch("http://localhost:8080/posts", {
            "method": "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(request)
        }).then((response) => {
            if (!response.ok) {
                throw new Error(`Error: ${response.statusText}`);
            }
            return response.json();
        })
        .then((result) => callBack ? callBack(result) : result)
        .catch((err) =>  errorCallback ? errorCallback(err) : err)
    }

    const createThread = (
        user: IProfile,
        message: string,
        callBack?: (result: string) => void,
        errorCallback?: (reason: string) => void
    ) => {

        const request: ICreateThreadRequest = {
            authorId: user.id,
            content: message
        }

        fetch("http://localhost:8080/posts", {
            "method": "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(request)
        }).then((response) => {
            if (!response.ok) {
                throw new Error(`Error: ${response.statusText}`);
            }
            return response.json();
        })
        .then((result) => callBack ? callBack(result) : result)
        .catch((err) =>  errorCallback ? errorCallback(err) : err)
    }


    return {
        reply: replyToThread,
        createThread: createThread
    }
}