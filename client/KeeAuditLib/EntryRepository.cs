using KeeAuditLib.Specifications;
using KeePassLib;
using System.Collections.Generic;
using System.Linq;

namespace KeeAuditLib
{
    public class EntryRepository
    {
        private readonly PwDatabase _database;

        public EntryRepository(PwDatabase database)
        {
            _database = database;
        }

        public IReadOnlyList<PwEntry> Find(Specification<PwEntry> specification)
        {
            var predicate = specification.ToExpression().Compile();
            return _database.RootGroup.GetEntries(true).Where(predicate).ToList();
        }

        public IReadOnlyList<PwEntry> GetAllEntries()
        {
            return _database.RootGroup.GetEntries(true).ToList();
        }

        public PwEntry GetEntryByUuid(PwUuid uuid)
        {
            return _database.RootGroup.GetEntries(true).FirstOrDefault(e => e.Uuid.Equals(uuid));
        }

        public int Count(Specification<PwEntry> specification)
        {
            var predicate = specification.ToExpression().Compile();
            return _database.RootGroup.GetEntries(true).Count(predicate);
        }
    }
}
