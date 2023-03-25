package db

import (
	"github.com/fairytale5571/punkz/internal/site"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *database) InsertWallet(data site.WalletDatabase) error {
	collection := db.mainDB.Collection("wallets")
	_, err := collection.InsertOne(db.ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (db *database) GetWallets() ([]site.WalletDatabase, error) {
	collection := db.mainDB.Collection("wallets")
	filter := bson.D{}
	cursor, err := collection.Find(db.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(db.ctx)

	var wallets []site.WalletDatabase
	for cursor.Next(db.ctx) {
		var wallet site.WalletDatabase
		if err := cursor.Decode(&wallet); err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(wallets) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return wallets, nil
}

func (db *database) GetWallet(userID string) (site.WalletDatabase, error) {
	collection := db.mainDB.Collection("wallets")
	var walletData site.WalletDatabase
	err := collection.FindOne(db.ctx, bson.M{
		"user_id": userID,
	}).Decode(&walletData)
	if err != nil {
		return site.WalletDatabase{}, err
	}
	return walletData, nil
}
