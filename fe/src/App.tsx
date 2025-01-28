import { useState } from 'react'
import dayjs from 'dayjs'
import axios from 'axios'
import { Input } from './Input'

function App() {
    const today = dayjs()
    const [day, setDay] = useState(today.date() + '')
    const [month, setMonth] = useState(today.month() + 1 + '')
    const [events, setEvents] = useState<string[]>([])
    const [loading, setLoading] = useState(false)

    const submit = async () => {
        setLoading(true)
        const endpoint =
            import.meta.env.VITE_BE || 'http://localhost:8080/list-event'

        const form = { day: +day, month: +month }

        const { data: resp } = await axios.post(endpoint, form)

        const events: string[] = resp.data.events
        setEvents(events)
        console.log(events)
        setLoading(false)
    }

    return (
        <div className="md:max-w-md md:mx-auto min-h-screen md:px-0 p-4 py-8 text-slate-100">
            <div className="px-4 py-8 shadow-lg border border-black">
                <Input id="day" label="Day" value={day} onChange={setDay} />
                <div className="pb-4" />
                <Input
                    id="month"
                    label="Month"
                    value={month}
                    onChange={setMonth}
                />
                <div className="pb-16" />
                <div className="-mx-4 -my-8">
                    <button
                        disabled={loading}
                        onClick={submit}
                        className="px-3 py-2 bg-purple-500 w-full text-center cursor-pointer"
                    >
                        {loading ? 'Loading...' : 'Submit'}
                    </button>
                </div>
            </div>

            <div className="flex flex-col py-8">
                {events.map((e, i) => (
                    <div key={i} className="p-2 border mb-2 border-black">
                        {e}
                    </div>
                ))}
            </div>
        </div>
    )
}

export default App
