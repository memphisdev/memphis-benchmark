const memphis = require('memphis-dev');

(async function () {
    try {
        await memphis.connect({
            host: 'localhost',
            username: 'benchmark',
            connectionToken: 'AhrnKXq5UZSBafNNXGnj' //mem user add -u benchmark --type application
        });

        const producer = await memphis.producer({
            stationName: 'benchmark',
            producerName: 'demo_producer'
        });

        const promises = [];
        for (let index = 0; index < 100; index++){
            promises.push(
                producer.produce({
                    message: Buffer.from(`test`)
                })
            );
            console.log("Message sent");
        }
        await Promise.all(promises);
        console.log("All messages sent");
        memphis.close();
    } catch (ex) {
        console.log(ex);
        memphis.close();
    }
})();
