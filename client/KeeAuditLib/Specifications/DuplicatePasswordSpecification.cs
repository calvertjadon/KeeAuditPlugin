using KeePassLib;
using System;
using System.Linq;
using System.Linq.Expressions;

namespace KeeAuditLib.Specifications
{
    public class DuplicatePasswordSpecification : Specification<PwEntry>
    {
        private readonly EntryRepository _repository;

        public DuplicatePasswordSpecification(EntryRepository repository)
        {
            _repository = repository;
        }

        private bool HasDuplicatePassword(PwEntry entry)
        {
            var password = entry.Strings.Get("Password").ReadString();
            return _repository.GetAllEntries()
                .Where(e => e.Uuid != entry.Uuid)
                .Any(e => e.Strings.Get("Password").ReadString() == password);
        }

        public override Expression<Func<PwEntry, bool>> ToExpression()
        {
            return entry => HasDuplicatePassword(entry);
        }
    }
}
