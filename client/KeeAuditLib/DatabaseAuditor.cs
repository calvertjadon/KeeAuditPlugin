using KeeAuditLib.Specifications;
using KeePassLib;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace KeeAuditLib
{
    public class DatabaseAuditor
    {
        private readonly EntryRepository _entryRepository;
        private readonly HashSet<Specification<PwEntry>> _specifications
            = new HashSet<Specification<PwEntry>>();

        public IEnumerable<Specification<PwEntry>> Specifications => _specifications;

        public DatabaseAuditor(EntryRepository entryRepository)
        {
            _entryRepository = entryRepository;
        }

        public void AddSpecification(Specification<PwEntry> specification)
        {
            _specifications.Add(specification);
        }

        public void RemoveSpecification(Specification<PwEntry> specification)
        {
            _specifications.Remove(specification);
        }


        public IEnumerable<ValidationResult> Validate()
        {
            var entries = _entryRepository.GetAllEntries();
            foreach (var entry in entries)
            {
                var failedSpecifications = _specifications.Where(
                    s => !s.IsSatisfiedBy(entry)
                ).ToList();
                yield return new ValidationResult(entry, failedSpecifications);
            }
        }
    }
}
