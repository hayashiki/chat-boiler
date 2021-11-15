import { NextPage } from 'next'
import { Message, useMessagePostedSubscription, useMessagesQuery, usePostMessageMutation } from "@/generated/graphql";
import { useRouter } from "next/router";
import { useCallback, useEffect, useState } from "react";

const Index: NextPage = () => {

    // const { roomId } = useRouter()
    const roomId = "0512cf5a-f822-4904-8117-aab8e2f9e4aa"

    const [messages, setMessages] = useState<Message[]>([]);
    const [inputValue, setInputValue] = useState('')
    const { loading, data, error } = useMessagesQuery({
        variables: {
            roomId: roomId
        }
    })
    const [createMessage] = usePostMessageMutation()

    useEffect(() => {
        if (data?.messages) setMessages(data.messages)
    }, [data?.messages])

    const { data: subData } = useMessagePostedSubscription({
      variables: {
        roomId: roomId,
      },
        onSubscriptionData: (data) => {
            console.log(data)
        }
    })

  useEffect(() => {
    if (subData?.messagePosted) setMessages(m => [...m, subData?.messagePosted])
  }, [subData?.messagePosted])

    const handleChange = useCallback((e: React.ChangeEvent<HTMLInputElement>) => {
        setInputValue(e.target.value)
    }, [])

  const handleClick = useCallback(async (e) => {
    e.preventDefault()
    await createMessage({ variables: { roomId: roomId, text: inputValue } })
  }, [inputValue, createMessage])

    if (data?.messages.length == 0) {
        return <div>No data found</div>
    }

    return <div>
        {messages.map((m, i) => <div key={i}>{m.text}</div>)}
        <input
          value={inputValue}
          onChange={handleChange}
        />
        <button onClick={handleClick}>Submit</button>
    </div>
}

export default Index;
