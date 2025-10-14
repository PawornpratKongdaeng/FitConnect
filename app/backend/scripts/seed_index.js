db = db.getSiblingDB('fitconnect')
db.gyms.createIndex({ location: '2dsphere' })